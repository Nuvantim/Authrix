package service

import (
	db "api/database"
	repo "api/internal/app/repository"
	req "api/internal/app/request"
	"api/pkgs/guards"
	rds "api/redis"

	ctx "context"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func SendOTP(email string) (string, error) {
	otp := guard.GenerateOTP()

	token := req.CreateOTP{
		Code:  otp,
		Email: email,
	}

	var key string = fmt.Sprintf("verif:%s", token.Code)
	if err := rds.SetData(key, token); err != nil {
		return "", err
	}

	fmt.Println(otp)

	// send otp via email
	if error := guard.SendOTP(token.Email, token.Code); error != nil {
		return "", error
	}

	return "otp send successfully", nil

}

func Register(regist req.Register) (string, error) {
	var data req.CreateOTP
	// search otp
	var value string = fmt.Sprintf("verif:%s", regist.Code)
	result, err := rds.GetData(value, data)
	if err != nil {
		return "", err
	}
	if result.Email == "" {
		return "", fmt.Errorf("OTP not found or expired")
	}

	// check email
	if result.Email != regist.Email {
		return "", fmt.Errorf("registration rejected")
	}

	var pass = guard.HashBycrypt(regist.Password) // Hashing Password
	// Regist New User
	createUser := repo.CreateUserParams{
		Name:     regist.Name,
		Email:    regist.Email,
		Password: string(pass),
	}

	// create user_account
	id_user, err := db.Queries.CreateUser(ctx.Background(), createUser)
	if err != nil {
		return "", db.Fatal(err)
	}
	// Create a buffered channel to receive any error from the goroutine
	errChan := make(chan error, 1)

	// Run user creation and OTP deletion in a separate goroutine
	go func() {
		// Create the user_profile
		if err := db.Queries.CreateProfile(ctx.Background(), id_user); err != nil {
			errChan <- db.Fatal(err) // Send error if user creation fails
			return
		}

		// Delete the used OTP
		if err := rds.DelData(value); err != nil {
			errChan <- err // Send error if OTP deletion fails
			return
		}

		// Both operations succeeded
		errChan <- nil
	}()

	// Wait for the result from the goroutine
	if err := <-errChan; err != nil {
		return "", err
	}

	return "your account has been created, please login", nil
}

func Login(login req.Login) (string, string, error) {
	// Find data account
	data, err := db.Queries.FindEmail(ctx.Background(), login.Email)
	if err != nil {
		return "", "", errors.New("account not found")
	}
	//compared password
	if err := bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(login.Password)); err != nil {
		return "", "", errors.New("password incorect")
	}

	role, err := db.Queries.AllRoleClient(ctx.Background(), data.ID)
	if err != nil {
		return "", "", db.Fatal(err)
	}
	// Input jwt data
	// Create access token and refresh token
	accessToken, err := guard.CreateToken(data.ID, data.Email, role)
	if err != nil {
		return "", "", db.Fatal(err)
	}

	refreshToken, err := guard.CreateRefreshToken(data.ID, data.Email)
	if err != nil {
		return "", "", db.Fatal(err)
	}

	return accessToken, refreshToken, nil

}

func ResetPassword(pass req.ResetPassword) (string, error) {
	// Check Code Otp
	var data req.CreateOTP
	var value string = fmt.Sprintf("verif:%s", pass.Code)
	result, err := rds.GetData(value, data)
	if err != nil {
		return "", err
	}
	if result.Email == "" {
		return "", fmt.Errorf("OTP not found or expired")
	}

	// Check Email User Account
	email_search, err := db.Queries.FindEmail(ctx.Background(), result.Email)
	if err != nil {
		return "", errors.New("email not found")
	}

	// Check if Password is same
	if pass.RetypePassword != pass.Password {
		return "", errors.New("password incorect")
	}

	// UpdatePassword
	psw := guard.HashBycrypt(pass.Password) //Hashing Password
	resetPassword := repo.ResetPasswordParams{
		Email:    email_search.Email,
		Password: string(psw),
	}
	// Create a buffered channel to receive any error from the goroutine
	errChan := make(chan error, 1)

	// Run database operations in a separate goroutine
	go func() {

		// Try to update the password
		if err := db.Queries.ResetPassword(ctx.Background(), resetPassword); err != nil {
			errChan <- db.Fatal(err)
			return
		}

		// Delete OTP code by email
		if err := rds.DelData(value); err != nil {
			errChan <- err
			return
		}
		errChan <- nil
	}()

	if err := <-errChan; err != nil {
		return "", err
	}

	return "reset password successfully", nil
}
