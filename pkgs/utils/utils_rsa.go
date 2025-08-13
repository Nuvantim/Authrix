package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"log"
	"os"
	"path/filepath"
)

// rsaKeyPath is the directory where the RSA keys are stored.
const rsaKeyPath = "./"

// GenRSA checks for the existence of RSA key files and generates a new key pair if they are missing.
func GenRSA() {
	privateKeyPath := filepath.Join(rsaKeyPath, "private.pem")
	publicKeyPath := filepath.Join(rsaKeyPath, "public.pem")

	_, errPublic := os.Stat(publicKeyPath)
	_, errPrivate := os.Stat(privateKeyPath)

	if os.IsNotExist(errPublic) || os.IsNotExist(errPrivate) {
		print("Generate RSA key....")
		privateKey, publicKey, err := generateRSAKeyPair(4096)
		if err != nil {
			log.Println("Gagal menghasilkan kunci RSA:", err)
			return
		}

		if err := savePEMKey(privateKeyPath, privateKey); err != nil {
			log.Println("Gagal menyimpan private key:", err)
		}

		if err := savePublicPEMKey(publicKeyPath, publicKey); err != nil {
			log.Println("Gagal menyimpan public key:", err)
		}
	}
}

// generateRSAKeyPair generates a new RSA key pair with the specified number of bits.
func generateRSAKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}
	return privateKey, &privateKey.PublicKey, nil
}

// savePEMKey saves an RSA private key to a file in PEM format.
func savePEMKey(filename string, key *rsa.PrivateKey) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	privateKeyBytes := x509.MarshalPKCS1PrivateKey(key)
	privateKeyPEM := &pem.Block{
		Type: "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}
	return pem.Encode(file, privateKeyPEM)
}

// savePublicPEMKey saves an RSA public key to a file in PEM format.
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
		Type: "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}
	return pem.Encode(file, publicKeyPEM)
}
