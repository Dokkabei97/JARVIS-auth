package domain

import (
	"github.com/golang-jwt/jwt/v5"
	"log"
	"time"
)

type JwtToken struct {
	Id           int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	UserId       int64     `json:"userId" gorm:"not null;column:user_id"`
	AccessToken  string    `json:"accessToken" gorm:"not null;column:access_token"`
	RefreshToken string    `json:"refreshToken" gorm:"not null;column:refresh_token"`
	CreatedAt    time.Time `json:"createdAt" gorm:"autoCreateTime;not null;column:created_at"`
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

func ValidateToken(jwtToken string, secretKey []byte) bool {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		log.Fatalf("[ERROR] ValidateToken : %s\n", err)
	}

	if token.Valid {
		return true
	} else {
		return false
	}
}
