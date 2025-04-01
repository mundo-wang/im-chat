package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/mundo-wang/wtool/wlog"
	"im-chat/code"
	"time"
)

type UserClaims struct {
	jwt.RegisteredClaims
	UserID   int    `json:"userId"`
	UserName string `json:"userName"`
	Phone    string `json:"phone"`
}

func GenerateJwtToken(userId int, userName, phone string) (string, error) {
	jwtKey := Config.Jwt.SecretKey
	claims := &UserClaims{
		UserID:   userId,
		UserName: userName,
		Phone:    phone,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        uuid.NewString(),                                   // 设置jti，标识每个jwt
			Issuer:    "Mundo",                                            // 设置发行者iss
			IssuedAt:  jwt.NewNumericDate(time.Now()),                     // 设置创建时间iat
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(72 * time.Hour)), // 设置过期时间exp为3天
		},
	}
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	jwtToken, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		wlog.Error("call token.SignedString failed").Err(err).Field("claims", claims).Log()
		return "", err
	}
	return jwtToken, nil
}

func ParseJwtToken(jwtToken string) (*UserClaims, error) {
	jwtKey := Config.Jwt.SecretKey
	token, err := jwt.ParseWithClaims(jwtToken, &UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil // 返回用于生成jwtToken时的密钥
		})
	if err != nil {
		wlog.Error("call jwt.ParseWithClaims failed").Err(err).Field("jwtToken", jwtToken).Log()
		return nil, err
	}
	claims, ok := token.Claims.(*UserClaims)
	if ok && token.Valid {
		return claims, nil
	} else {
		return nil, code.JwtTokenInvalid
	}
}
