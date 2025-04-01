package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestGenerateSalt(t *testing.T) {
	data := "Hello, World!"
	salt, err := GenerateSalt(32)
	if err != nil {
		t.Errorf("call GenerateSalt failed, error: %v", err)
	}
	dataWithSalt := data + salt
	hasher := sha256.New()
	hasher.Write([]byte(dataWithSalt))
	hash := hasher.Sum(nil)
	hashString := hex.EncodeToString(hash)
	fmt.Println(hashString)
}
