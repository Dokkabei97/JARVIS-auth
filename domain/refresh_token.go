package domain

import (
	"github.com/golang-jwt/jwt/v5"
	"log"
)

type RefreshToken struct {
	UserId int64 `json:"userId"`
	jwt.RegisteredClaims
}

func GenerateRefreshToken(userId int64, secretKey []byte) string {
	refreshToken := RefreshToken{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: nil,
			Issuer:    "",
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshToken).
		SignedString(secretKey)

	if err != nil {
		log.Fatalf("[ERROR] GenerateRefreshToken : %s\n", err)
	}

	return token
}
func ValidateRefreshToken(refreshToken string, secretKey []byte) bool {
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		log.Fatalf("[ERROR] ValidateRefreshToken : %s\n", err)
	}

	if token.Valid {
		return true
	} else {
		return false
	}
}
