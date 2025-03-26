package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/mundo-wang/wtool/wlog"
	"testing"
)

func TestGenerateSalt(t *testing.T) {
	data := "Hello, World!"
	salt, err := GenerateSalt(32)
	if err != nil {
		wlog.Error("call GenerateSalt failed").Err(err).Log()
	}
	dataWithSalt := data + salt
	hasher := sha256.New()
	hasher.Write([]byte(dataWithSalt))
	hash := hasher.Sum(nil)
	hashString := hex.EncodeToString(hash)
	fmt.Println(hashString)
}
