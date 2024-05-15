package rsa

import (
	"crypto/rand"
	RSAlib "crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
)

func (r *RSA) Encode(publicKeyStr, plaintext string) (encodedText string, err error) {
	publicKeyBlock, _ := pem.Decode([]byte(publicKeyStr))
	if publicKeyBlock == nil {
		return "", fmt.Errorf("public key decode error")
	}
	publicKeyObj, err := x509.ParsePKIXPublicKey(publicKeyBlock.Bytes)
	if err != nil {
		return "", fmt.Errorf("public key parsing error: %v", err)
	}
	rsaPub, ok := publicKeyObj.(*RSAlib.PublicKey)
	if !ok {
		return "", fmt.Errorf("public key casting error")
	}

	ciphertext, err := RSAlib.EncryptPKCS1v15(rand.Reader, rsaPub, []byte(plaintext))
	if err != nil {
		return "", err
	}

	encodedText = base64.StdEncoding.EncodeToString(ciphertext)
	return encodedText, nil
}
