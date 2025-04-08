package utils

import (
	"crypto/rand"
	"fmt"
	"github.com/mundo-wang/wtool/wlog"
	"math/big"
	"strings"
)

func GenerateRandomDigits(n int) (string, error) {
	if n <= 5 {
		return "", fmt.Errorf("生成数字的长度必须大于5")
	}
	var sb strings.Builder
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(10))
		if err != nil {
			wlog.Error("call rand.Int failed").Log()
			return "", err
		}
		if i == 0 && num.Int64() == 0 {
			i--
			continue
		}
		sb.WriteString(fmt.Sprintf("%d", num.Int64()))
	}
	return sb.String(), nil
}
