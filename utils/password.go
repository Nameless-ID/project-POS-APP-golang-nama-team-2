package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
)

func GeneratePassword(password string, key *rsa.PrivateKey) (string, error) {
	hash := sha256.New()
	encryptedBytes, err := rsa.EncryptOAEP(hash, rand.Reader, &key.PublicKey, []byte(password), nil)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(encryptedBytes), nil
}
