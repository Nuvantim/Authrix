package service

import (
	db "api/database"
	repo "api/internal/repository"
	req "api/internal/request"
	"api/pkgs/utils"

	ctx "context"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func SendOTP(email string) (string, error) {
	otp := utils.GenerateOTP()

	token := repo.CreateOTPParams{
		Code:  otp,
		Email: email,
	}

	if err := db.Queries.CreateOTP(ctx.Background(), token); err != nil {
		return "", err
	}
	return "OTP send successfully", nil

}

func Register(regist req.Register) (string, error) {
	// search otp code & email
	searchOtp := repo.FindOtpByEmailParams{
		Code:  regist.Code,
		Email: regist.Email,
	}

	otpSearch, err := db.Queries.FindOtpByEmail(ctx.Background(), searchOtp)
	if err != nil {
		return "", err
	}

	if otpSearch == 0 {
		return "", errors.New("OTP not found or invalid")
	}

	var pass = utils.HashBycrypt(regist.Password) // Hashing Password
	// Regist New User
	createUser := repo.CreateUserParams{
		Name:     regist.Name,
		Email:    regist.Email,
		Password: string(pass),
	}

	// Create a buffered channel to receive any error from the goroutine
	errChan := make(chan error, 1)

	// Run user creation and OTP deletion in a separate goroutine
	go func() {
		// Create the user
		if err := db.Queries.CreateUser(ctx.Background(), createUser); err != nil {
			errChan <- err // Send error if user creation fails
			return
		}

		// Delete the used OTP
		if err := db.Queries.DeleteOTP(ctx.Background(), createUser.Email); err != nil {
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

	return "Your account has been created, please login", nil
}

func Login(login req.Login) (string, error) {
	// Find data account
	data, err := db.Queries.FindEmail(ctx.Background(), login.Email)
	if err != nil {
		return "", errors.New("Account Not Found")
	}
	//compared password
	if err := bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(login.Password)); err != nil {
		return "", errors.New("Password Incorect")
	}

	// Create JWT
	/* coming soon*/

	return "login success", nil

}

func ResetPassword(pass req.ResetPassword) (string, error) {

	// Check Code Otp
	otp_search, err := db.Queries.FindOTP(ctx.Background(), pass.Code)
	if err != nil {
		return "", errors.New("OTP code not found")
	}

	// Check Email User Account
	email_search, err := db.Queries.FindEmail(ctx.Background(), otp_search.Email)
	if err != nil {
		return "", errors.New("Email not found")
	}

	// Check if Password is same
	if pass.RetypePassword != pass.Password {
		return "", errors.New("Password incorect")
	}

	// UpdatePassword
	psw := utils.HashBycrypt(pass.Password) //Hashing Password
	resetPassword := repo.ResetPasswordParams{
		Email:    email_search.Email,
		Password: string(psw),
	}
	// Create a buffered channel to receive any error from the goroutine
	errChan := make(chan error, 1)

	// Run database operations in a separate goroutine
	go func() {

		// Try to update the password
		if err := db.Queries.ResetPassword(ctx.Background(), ResetPassword); err != nil {
			errChan <- err
			return
		}

		// Delete OTP code by email
		if err := db.Queries.DeleteOTP(ctx.Background(), email_search.Email); err != nil {
			errChan <- err
			return
		}
		errChan <- nil
	}()

	if err := <-errChan; err != nil {
		return "", err
	}

	return "Reset password successfully", nil
}
