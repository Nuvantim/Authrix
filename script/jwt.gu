package utils

import (
	"api/internal/database"
	"api/internal/domain/models"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

// get api key from .env
var jwtSecret = []byte(os.Getenv("API_KEY"))
var refreshSecret = []byte(os.Getenv("REFRESH_KEY"))

// Claims untuk Access Token
type Claims struct {
	UserID uint          `json:"user_id"`
	Email  string        `json:"email"`
	Roles  []models.Role `json:"roles"`
	jwt.RegisteredClaims
}

// Claims untuk Refresh Token
type RefreshClaims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

// Buat Access Token
func CreateToken(userID uint, Email string, Roles []models.Role) (string, error) {
	now := time.Now()

	claims := Claims{
		UserID: userID,
		Email:  Email,
		Roles:  Roles,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),                       // Add issued at
			ExpiresAt: jwt.NewNumericDate(now.Add(2 * time.Hour)), // Token valid for 2 hours
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// Buat Refresh Token
func CreateRefreshToken(userID uint, email string) (string, error) {
	now := time.Now()

	claims := RefreshClaims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // Refresh token berlaku 30 hari
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(refreshSecret)
}

func AutoRefressToken(UserID uint) (string, error) {
	var user models.User
	database.DB.Preload("Roles").Preload("Roles.Permissions").Take(&user, UserID)

	now := time.Now()

	claims := Claims{
		UserID: user.ID,
		Email:  user.Email,
		Roles:  user.Roles,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),                    // Add issued at
			ExpiresAt: jwt.NewNumericDate(now.Add(2 * time.Hour)), // Token valid for 2 hours
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
