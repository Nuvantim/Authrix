package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"log"
	"os"
	"sync"
)

func GenerateRSA(privateKeyFile, publicKeyFile string, bits int) error {
	// Check if both key files already exist.
	_, privateErr := os.Stat(privateKeyFile)
	_, publicErr := os.Stat(publicKeyFile)

	// If both files are found, no need to generate new keys.
	if privateErr == nil && publicErr == nil {
		return nil
	}

	// If either file is missing, generate new keys.
	if os.IsNotExist(privateErr) || os.IsNotExist(publicErr) {
		log.Println("Generating RS512 key pair...")

		var wg sync.WaitGroup
		wg.Add(1)

		go func() {
			defer wg.Done()

			// Generate the new RSA key pair.
			privateKey, publicKey, err := generateRSAKeyPair(bits)
			if err != nil {
				log.Println("Failed to generate RSA keys:", err)
				return
			}

			// Save the private key.
			if err := savePEMKey(privateKeyFile, privateKey); err != nil {
				log.Println("Failed to save private key:", err)
				return
			}

			// Save the public key.
			if err := savePublicPEMKey(publicKeyFile, publicKey); err != nil {
				log.Println("Failed to save public key:", err)
				return
			}

		}()

		// Wait for the goroutine to complete.
		wg.Wait()
		return nil
	}

	// Handle other potential errors (e.g., permission issues).
	log.Println("Error checking key files:")
	log.Println("Private key error:", privateErr)
	log.Println("Public key error:", publicErr)
	return privateErr
}

// generateRSAKeyPair generates an RSA key pair (private and public).
func generateRSAKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}
	return privateKey, &privateKey.PublicKey, nil
}

// savePEMKey saves the private key to a file in PEM format.
func savePEMKey(filename string, key *rsa.PrivateKey) error {
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

// savePublicPEMKey saves the public key to a file in PEM format.
func savePublicPEMKey(filename string, pubkey *rsa.PublicKey) error {
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
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKeyBytes,
	}
	return pem.Encode(file, publicKeyPEM)
}
