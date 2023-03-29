package domain

import (
	"github.com/golang-jwt/jwt/v5"
	"go/token"
)

type RefreshToken struct {
	UserId int64 `json:"userId"`
	jwt.RegisteredClaims
}

func GenerateRefreshToken(userId int64) *token.Token {
	return nil
}
func ValidateRefreshToken(refreshToken string) {

}
