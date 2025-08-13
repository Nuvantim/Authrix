package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// rsaKeyPath is the directory where the RSA keys are stored.
// Change this path if you want to store keys in a different folder.
const rsaKeyPath = "./"

// GenRSA checks if RSA key files exist; if not, it generates a new key pair.
func GenRSA() {
	privateKeyPath := filepath.Join(rsaKeyPath, "private.pem")
	publicKeyPath := filepath.Join(rsaKeyPath, "public.pem")

	_, errPublic := os.Stat(publicKeyPath)
	_, errPrivate := os.Stat(privateKeyPath)

	// If either public or private key file does not exist, generate new keys
	if os.IsNotExist(errPublic) || os.IsNotExist(errPrivate) {
		fmt.Println("Generating RSA key pair...")
		privateKey, publicKey, err := generateRSAKeyPair(4096)
		if err != nil {
			log.Println("Failed to generate RSA keys:", err)
			return
		}

		if err := savePEMKey(privateKeyPath, privateKey); err != nil {
			log.Println("Failed to save private key:", err)
		}

		if err := savePublicPEMKey(publicKeyPath, publicKey); err != nil {
			log.Println("Failed to save public key:", err)
		}
	}
}

// generateRSAKeyPair creates a new RSA key pair with the specified bit size.
func generateRSAKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}
	return privateKey, &privateKey.PublicKey, nil
}

// savePEMKey saves the RSA private key to a PEM-formatted file.
func savePEMKey(filename string, key *rsa.PrivateKey) error {
	// Ensure the file path is safe and within the allowed directory
	if !isSafePath(filename, rsaKeyPath) {
		return fmt.Errorf("invalid file path: %s", filename)
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	privateKeyBytes := x509.MarshalPKCS1PrivateKey(key)
	privateKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}
	return pem.Encode(file, privateKeyPEM)
}

// savePublicPEMKey saves the RSA public key to a PEM-formatted file.
func savePublicPEMKey(filename string, pubkey *rsa.PublicKey) error {
	// Ensure the file path is safe and within the allowed directory
	if !isSafePath(filename, rsaKeyPath) {
		return fmt.Errorf("invalid file path: %s", filename)
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	publicKeyBytes, err := x509.MarshalPKIXPublicKey(pubkey)
	if err != nil {
		return err
	}

	publicKeyPEM := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}
	return pem.Encode(file, publicKeyPEM)
}

// isSafePath ensures the final absolute file path remains within the given base directory.
// This prevents path traversal attacks (e.g., using ../../ to escape the base directory).
func isSafePath(filePath, baseDir string) bool {
	absBase, err1 := filepath.Abs(baseDir)
	absTarget, err2 := filepath.Abs(filePath)
	if err1 != nil || err2 != nil {
		return false
	}
	return filepath.Dir(absTarget) == absBase
}
