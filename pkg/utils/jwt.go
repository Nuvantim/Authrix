package utils

import (
	"api/internal/database"
	"api/internal/domain/models"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"io/ioutil"
	"log"
	"time"
)

var (
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
)

// Load Private Key File
func LoadPrivateKey() (*rsa.PrivateKey, error) {
	keyBytes, err := ioutil.ReadFile("private.pem")
	if err != nil {
		return nil, err
	}

	block, rest := pem.Decode(keyBytes)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, errors.New("invalid private key format")
	}

	// Handle any remaining data after decoding
	if len(rest) > 0 {
		return nil, errors.New("extra data found after private key")
	}

	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

// Load Public Key File
func LoadPublicKey() (*rsa.PublicKey, error) {
	keyBytes, err := ioutil.ReadFile("public.pem")
	if err != nil {
		return nil, err
	}

	block, rest := pem.Decode(keyBytes)
	if block == nil || block.Type != "RSA PUBLIC KEY" {
		return nil, errors.New("invalid public key format")
	}

	// Handle any remaining data after decoding
	if len(rest) > 0 {
		return nil, errors.New("extra data found after public key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	rsaPubKey, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("invalid type for RSA public key")
	}

	return rsaPubKey, nil
}

// Initialize the keys
func init() {
	var err error
	privateKey, err = LoadPrivateKey()
	if err != nil {
		log.Fatalf("Error loading private key: %v", err)
	}

	publicKey, err = LoadPublicKey()
	if err != nil {
		log.Fatalf("Error loading public key: %v", err)
	}
}

// Buat Access Token
func CreateToken(userID uint, email string, roles []models.Role) (string, error) {
	// Ensure the private key is not nil before using it
	if privateKey == nil {
		return "", errors.New("private key is nil")
	}

	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"roles":   roles,
		"exp":     time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS512, claims)

	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Buat Refresh Token
func CreateRefreshToken(userID uint, email string) (string, error) {
	// Ensure the private key is not nil before using it
	if privateKey == nil {
		return "", errors.New("private key is nil")
	}

	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS512, claims)

	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// AutoRefreshToken
func AutoRefreshToken(userID uint) (string, error) {
	var user models.User
	// Preloading roles and permissions
	if err := database.DB.Preload("Roles").Preload("Roles.Permissions").Take(&user, userID).Error; err != nil {
		return "", err
	}

	token, err := CreateToken(user.ID, user.Email, user.Roles)
	if err != nil {
		return "", err
	}

	return token, nil
}
