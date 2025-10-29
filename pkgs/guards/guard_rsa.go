package guard

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var rsaKeyPath = "./screet-key"

// GenRSA generates an RSA key pair if the key files don't exist.
func GenRSA() {
	// Create folder if not exists
	if err := os.MkdirAll(rsaKeyPath, 0700); err != nil {
		log.Println("Failed to create directory:", err)
		return
	}

	privateKeyPath := filepath.Join(rsaKeyPath, "private.pem")
	publicKeyPath := filepath.Join(rsaKeyPath, "public.pem")

	_, errPublic := os.Stat(publicKeyPath)
	_, errPrivate := os.Stat(privateKeyPath)

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

// generateRSAKeyPair creates a new RSA private and public key pair.
func generateRSAKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	key, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}
	return key, &key.PublicKey, nil
}

// savePEMKey writes a private RSA key to a PEM file.
func savePEMKey(filename string, key *rsa.PrivateKey) error {
	if !isSafePath(filename, rsaKeyPath) {
		return fmt.Errorf("invalid file path: %s", filename)
	}

	fullpath := filepath.Join(rsaKeyPath, filepath.Base(filename))
	file, err := os.Create(fullpath)
	if err != nil {
		return err
	}
	defer file.Close()

	block := &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)}
	return pem.Encode(file, block)
}

// savePublicPEMKey writes a public RSA key to a PEM file.
func savePublicPEMKey(filename string, pub *rsa.PublicKey) error {
	if !isSafePath(filename, rsaKeyPath) {
		return fmt.Errorf("invalid file path: %s", filename)
	}

	fullpath := filepath.Join(rsaKeyPath, filepath.Base(filename))
	file, err := os.Create(fullpath)
	if err != nil {
		return err
	}
	defer file.Close()

	bytes, err := x509.MarshalPKIXPublicKey(pub)
	if err != nil {
		return err
	}
	block := &pem.Block{Type: "PUBLIC KEY", Bytes: bytes}
	return pem.Encode(file, block)
}

// isSafePath ensures the file path stays within the allowed base directory.
func isSafePath(filePath, baseDir string) bool {
	base, err1 := filepath.Abs(baseDir)
	target, err2 := filepath.Abs(filePath)
	if err1 != nil || err2 != nil {
		return false
	}
	if !strings.HasSuffix(base, string(os.PathSeparator)) {
		base += string(os.PathSeparator)
	}
	return strings.HasPrefix(target, base)
}
