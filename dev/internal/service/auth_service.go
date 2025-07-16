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

	_, err := db.Queries.CreateOTP(ctx.Background(), token)
	if err != nil {
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

	// check otp validation
	if otpSearch == 0 {
		return "", errors.New("OTP not found or invalid")
	}

	// Hashing Password
	var pass = utils.HashBycrypt(regist.Password)
	// Regist New User
	createUser := repo.CreateUserParams{
		Name:     regist.Name,
		Email:    regist.Email,
		Password: string(pass),
	}

	_, err = db.Queries.CreateUser(ctx.Background(), createUser)
	if err != nil {
		return "", err
	}

	return "Your account has been created, please login", nil
}

