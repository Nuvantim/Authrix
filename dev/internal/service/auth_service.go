package service

import (
	db "api/database"
	repo "api/internal/repository"
	req "api/internal/request"
	"api/pkgs/utils"
	ctx "context"
	"errors"
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

	if err := db.Queries.CreateUser(ctx.Background(), createUser); err != nil {
		return "", err
	}

	if err := db.Queries.DeleteOTP(ctx.Background(), createUser.Email); err != nil {
		return "", err
	}

	return "Your account has been created, please login", nil
}

func Login(login req.Login) (string, error) {
	psw := utils.HashBycrypt(login.Password) //Hashing Password
	// Find data account
	login_user := repo.LoginUserParams{
		Email:    login.Email,
		Password: string(psw),
	}
	data, err := db.Queries.LoginUser(ctx.Background(), login_user)
	if err != nil {
		return "", err
	}
	// check data if not found
	if data.ID == 0 {
		return "", errors.New("Account Not Found")
	}

	// Create JWT
	/* coming soon*/

	return "jwt", nil

}

func UpdatePassword(pass req.UpdatePassword) (string, error) {

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
	updatePassword := repo.UpdatePasswordParams{
		Email:    email_search.Email,
		Password: string(psw),
	}

	if err := db.Queries.UpdatePassword(ctx.Background(), updatePassword); err != nil {
		return "", err
	}

	if err := db.Queries.DeleteOTP(ctx.Background(), email_search.Email); err != nil {
		return "", err
	}

	return "Password successfully updated", nil
}
