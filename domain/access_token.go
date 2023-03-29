package domain

import (
	"github.com/golang-jwt/jwt/v5"
	"go/token"
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

func GenerateAccessToken(userInfo *UserInfo) *token.Token {
	return nil
}

func ValidateAccessToken(accessToken string) {

}
