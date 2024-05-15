package rsa

import (
	"crypto/rand"
	RSAlib "crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
)

func Decode(privateKeyStr, encodedText string) (decodedText string, err error) {
	privateKeyBlock, _ := pem.Decode([]byte(privateKeyStr))
	if privateKeyBlock == nil {
		return "", fmt.Errorf("private key decode error")
	}
	privateKeyObj, err := x509.ParsePKCS1PrivateKey(privateKeyBlock.Bytes)
	if err != nil {
		return "", fmt.Errorf("private key parsing error: %v", err)
	}

	ciphertext, err := base64.StdEncoding.DecodeString(encodedText)
	if err != nil {
		return "", err
	}

	decryptedBytes, err := RSAlib.DecryptPKCS1v15(rand.Reader, privateKeyObj, ciphertext)
	if err != nil {
		return "", err
	}

	decodedText = string(decryptedBytes)
	return decodedText, nil
}
