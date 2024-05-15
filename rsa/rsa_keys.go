package rsa

import (
	"crypto/rand"
	RSAlib "crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

func GenerateKey() (publicKey string, privateKey string, err error) {
	privateKeyObj, err := RSAlib.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return "", "", err
	}

	// تبدیل کلید خصوصی به PEM
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKeyObj)
	privateKeyBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}
	privateKey = string(pem.EncodeToMemory(privateKeyBlock))

	// تبدیل کلید عمومی به PEM
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(&privateKeyObj.PublicKey)
	if err != nil {
		return "", "", fmt.Errorf("unable to marshal public key: %v", err)
	}
	publicKeyBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}
	publicKey = string(pem.EncodeToMemory(publicKeyBlock))

	return publicKey, privateKey, nil
}
