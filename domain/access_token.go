package domain

import (
	"github.com/golang-jwt/jwt/v5"
	"log"
)

type UserInfo struct {
	UserId  int64  `json:"userId"`
	Email   string `json:"email"`
	IsAdmin bool   `json:"isAdmin"`
	IsBlock bool   `json:"isBlock"`
}

type AccessToken struct {
	UserInfo UserInfo `json:"userInfo"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(userInfo UserInfo, secretKey []byte) string {
	accessToken := AccessToken{
		UserInfo: userInfo,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: nil,
			Issuer:    "nil",
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessToken).
		SignedString(secretKey)

	if err != nil {
		log.Fatalf("[ERROR] GenerateAccessToken : %s\n", err)
	}

	return token
}

func ValidateAccessToken(accessToken string, secretKey []byte) bool {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		log.Fatalf("[ERROR] ValidateAccessToken : %s\n", err)
	}

	if token.Valid {
		return true
	} else {
		return false
	}
}
