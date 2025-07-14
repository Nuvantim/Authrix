package service

import (
	db "api/database"
	repo "api/internal/repository"
	"context"
	"api/pkgs/utils"
)

func SendOTP(email string) (string, error) {
	_,otp := utils.GenerateOTP()

	token := repo.CreateOTPParams{
		Code:  otp,
		Email: email,
	}

	_, err := db.Queries.CreateOTP(context.Background(), token)
	if err != nil {
		return "", err
	}
	return "OTP send successfully", nil

}
