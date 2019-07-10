package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret []byte

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// 生成token字符串
func GenerateToken(username, password string) (string, error) {
	expireTime := time.Now().Add(3 * time.Hour) // token过期时间

	claims := Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "rehabilitation_prescription",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // token签名方法暂不知区别
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// 解析token字符串
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
