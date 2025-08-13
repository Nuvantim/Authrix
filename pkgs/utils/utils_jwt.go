package utils

import (
	db "api/database"
	repo "api/internal/app/repository"

	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
)

// Claims defines JWT access token claims
type Claims struct {
	UserID int32                   `json:"user_id"`
	Email  string                  `json:"email"`
	Roles  []repo.AllRoleClientRow `json:"roles,omitempty"`
	jwt.RegisteredClaims
}

// RefreshClaims defines JWT refresh token claims
type RefreshClaims struct {
	UserID int32  `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

// loadKey reads and parses an RSA key file securely
func loadKey(filename string, isPrivate bool) (interface{}, error) {
	// Ensure the path is safe using utils_rsa.go
	if !IsSafePath(filename, RsaKeyPath) {
		return nil, fmt.Errorf("invalid file path: %s", filename)
	}

	keyBytes, err := os.ReadFile(filename) // #nosec G304, path sudah tervalidasi
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(keyBytes)
	if block == nil {
		return nil, errors.New("failed to decode PEM block containing the key")
	}

	if isPrivate {
		if block.Type != "RSA PRIVATE KEY" {
			return nil, errors.New("invalid private key format")
		}
		return x509.ParsePKCS1PrivateKey(block.Bytes)
	}

	if block.Type != "PUBLIC KEY" {
		return nil, errors.New("invalid public key format")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	rsaPubKey, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("parsed public key is not an RSA key")
	}

	return rsaPubKey, nil
}

// LoadPrivateKey loads the private RSA key from file
func LoadPrivateKey() (*rsa.PrivateKey, error) {
	key, err := loadKey(RsaKeyPath+"/private.pem", true)
	if err != nil {
		return nil, err
	}
	return key.(*rsa.PrivateKey), nil
}

// LoadPublicKey loads the public RSA key from file
func LoadPublicKey() (*rsa.PublicKey, error) {
	key, err := loadKey(RsaKeyPath+"/public.pem", false)
	if err != nil {
		return nil, err
	}
	return key.(*rsa.PublicKey), nil
}

// CheckRSA initializes the RSA keys
func CheckRSA() {
	privateKey, err := LoadPrivateKey()
	if err != nil {
		log.Fatalf("Failed to load private key: %v", err)
	}
	PrivateKey = privateKey

	publicKey, err := LoadPublicKey()
	if err != nil {
		log.Fatalf("Failed to load public key: %v", err)
	}
	PublicKey = publicKey
}

// CreateToken generates an access token
func CreateToken(id int32, email string, role []repo.AllRoleClientRow) (string, error) {
	if PrivateKey == nil {
		return "", errors.New("private key is nil")
	}

	now := time.Now().UTC()
	claims := Claims{
		UserID: id,
		Email:  email,
		Roles:  role,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(30 * time.Second)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(PrivateKey)
}

// CreateRefreshToken generates a refresh token
func CreateRefreshToken(id int32, email string) (string, error) {
	if PrivateKey == nil {
		return "", errors.New("private key is nil")
	}

	now := time.Now().UTC()
	claims := RefreshClaims{
		UserID: id,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(PrivateKey)
}

// AutoRefreshToken automatically refreshes the JWT token
func AutoRefreshToken(userID int32) (string, error) {
	user, err := db.Queries.GetClient(context.Background(), userID)
	if err != nil {
		return "", err
	}

	role, err := db.Queries.AllRoleClient(context.Background(), user.ID)
	if err != nil {
		return "", err
	}

	freshJwt, err := CreateToken(user.ID, user.Email, role)
	if err != nil {
		return "", err
	}
	return freshJwt, nil
}
