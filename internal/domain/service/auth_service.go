package service

import (
	"api/internal/database"
	"api/internal/domain/models"
	"api/pkg/utils"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Login Action
func Login(email, password string) (string, string, error) {
	// Find User in Database
	var user models.User
	err := database.DB.Where("email = ?", email).Preload("Roles").Preload("Roles.Permissions").Take(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "", "", errors.New("user not found")
	} else if err != nil {
		return "", "", err
	}

	// Compared Database
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", "", errors.New("invalid email or password")
	}

	// Create access token and refresh token
	accessToken, err := utils.CreateToken(user.ID, user.Email, user.Roles)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := utils.CreateRefreshToken(user.ID, user.Email)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
