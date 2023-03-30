package domain

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JwtToken struct {
	Id           int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	UserId       int64     `json:"userId" gorm:"not null;column:user_id"`
	AccessToken  string    `json:"accessToken" gorm:"not null;column:access_token"`
	RefreshToken string    `json:"refreshToken" gorm:"not null;column:refresh_token"`
	CreatedAt    time.Time `json:"createdAt" gorm:"autoCreateTime;not null;column:created_at"`
}

const ISSUER = "Deploy"

// generateToken JWT 토큰 생성
func generateToken(claims jwt.Claims, secretKey []byte) (string, error) {
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).
		SignedString(secretKey)
	if err != nil {
		return "", fmt.Errorf("[ERROR] GenerateToken : %s\n", err)
	}
	return token, nil
}

var loc, _ = time.LoadLocation("Asia/Seoul")

// GenerateAccessToken Access JWT 토큰 생성
func GenerateAccessToken(userInfo UserInfo, secretKey []byte) (string, error) {
	expirationTime := time.Now().In(loc).Add(time.Hour)

	accessToken := &AccessToken{
		UserInfo: userInfo,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    ISSUER,
		},
	}

	return generateToken(accessToken, secretKey)
}

// GenerateRefreshToken Refresh JWT 토큰 생성
func GenerateRefreshToken(userId int64, secretKey []byte) (string, error) {
	expirationTime := time.Now().In(loc).Add(time.Hour * 24 * 30)

	refreshToken := &RefreshToken{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    ISSUER,
		},
	}

	return generateToken(refreshToken, secretKey)
}

// ValidateToken JWT 토큰 검증
func ValidateToken(jwtToken string, secretKey []byte) (bool, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil || !token.Valid {
		return false, fmt.Errorf("[ERROR] ValidateToken : %s\n", err)
	}

	return true, nil
}
