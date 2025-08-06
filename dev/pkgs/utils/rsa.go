package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"log"
	"os"
)

func GenRSA() {
	// Check RSA file
	_, err_public := os.Stat("public.pem")
	_, err_private := os.Stat("private.pem")

	if os.IsNotExist(err_public) || os.IsNotExist(err_private) {
		print("Generate RSA key....")
		privateKey, publicKey, err := generateRSAKeyPair(4096)
		if err != nil {
			log.Println("Gagal menghasilkan kunci RSA:", err)
		}

		if err := savePEMKey("private.pem", privateKey); err != nil {
			log.Println("Gagal menyimpan private key:", err)
		}

		if err := savePublicPEMKey("public.pem", publicKey); err != nil {
			log.Println("Gagal menyimpan public key:", err)
		}
	}
}

func generateRSAKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}
	return privateKey, &privateKey.PublicKey, nil
}

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
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}
	return pem.Encode(file, publicKeyPEM)
}
