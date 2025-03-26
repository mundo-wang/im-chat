package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"github.com/mundo-wang/wtool/wlog"
)

func GenerateSalt(size int) (string, error) {
	salt := make([]byte, size)
	_, err := rand.Read(salt)
	if err != nil {
		wlog.Error("call rand.Read failed").Err(err).Log()
		return "", err
	}
	return base64.StdEncoding.EncodeToString(salt), nil
}

func GenerateSignature(data string) (string, string, error) {
	salt, err := GenerateSalt(32)
	if err != nil {
		wlog.Error("call GenerateSalt failed").Err(err).Log()
		return "", "", err
	}
	dataWithSalt := data + salt
	hasher := sha256.New()
	hasher.Write([]byte(dataWithSalt))
	hash := hasher.Sum(nil)
	signature := base64.StdEncoding.EncodeToString(hash)
	return signature, salt, nil
}

func VerifySignature(data, salt, receivedSignature string) bool {
	dataWithSalt := data + salt
	hasher := sha256.New()
	hasher.Write([]byte(dataWithSalt))
	hash := hasher.Sum(nil)
	calculatedSignature := base64.StdEncoding.EncodeToString(hash)
	return calculatedSignature == receivedSignature
}
