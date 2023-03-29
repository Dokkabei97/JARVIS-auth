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

func GenerateAccessToken(userInfo UserInfo) string {
	accessToken := AccessToken{
		UserInfo: userInfo,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: nil,
			Issuer:    "nil",
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessToken).
		SignedString("")

	if err != nil {
		log.Fatalf("[ERROR] GenerateAccessToken : %s\n", err)
	}

	return token
}

func ValidateAccessToken(accessToken string) {

}
