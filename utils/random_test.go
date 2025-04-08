package utils

import (
	"fmt"
	"testing"
)

func TestGenerateRandomDigits(t *testing.T) {
	digits, err := GenerateRandomDigits(10)
	if err != nil {
		t.Errorf("call GenerateRandomDigits failed, error: %v", err)
	}
	fmt.Println(digits)
}
