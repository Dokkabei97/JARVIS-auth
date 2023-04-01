package domain

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"is-deploy-auth/util"
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

var secretKey = util.GetSecretKey()

// generateToken JWT 토큰 생성
// 첫번째 리턴값 : 토큰
// 두번째 리턴값 : 에러
// 토큰 생성 실패시 에러 리턴
// 토큰 생성 성공시 토큰 리턴
func generateToken(claims jwt.Claims) (string, error) {
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).
		SignedString(secretKey)
	if err != nil {
		return "", fmt.Errorf("[ERROR] GenerateToken : %s\n", err)
	}
	return token, nil
}

var loc, _ = time.LoadLocation("Asia/Seoul")

// GenerateAccessToken Access JWT 토큰 생성
// 1시간간 유효
// 첫번째 리턴값 : 토큰
// 두번째 리턴값 : 에러
func GenerateAccessToken(userInfo UserInfo) (string, error) {
	expirationTime := time.Now().In(loc).Add(time.Hour)

	accessToken := &AccessToken{
		UserInfo: userInfo,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    ISSUER,
		},
	}

	return generateToken(accessToken)
}

// GenerateRefreshToken Refresh JWT 토큰 생성
// 30일간 유효
// 첫번째 리턴값 : 토큰
// 두번째 리턴값 : 에러
func GenerateRefreshToken(userId int64) (string, error) {
	expirationTime := time.Now().In(loc).Add(time.Hour * 24 * 30)

	refreshToken := &RefreshToken{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    ISSUER,
		},
	}

	return generateToken(refreshToken)
}

// ValidateToken JWT 토큰 검증
// 첫번째 리턴값 : 토큰 검증 여부
// 두번째 리턴값 : 에러
func ValidateToken(jwtToken string) (bool, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil || !token.Valid {
		return false, fmt.Errorf("[ERROR] ValidateToken : %s\n", err)
	}

	return true, nil
}

// ValidateAdmin 관리자 권한 검증
// 첫번째 리턴값 : 토큰 검증 여부
// 두번째 리턴값 : 관리자 권한 여부
// 세번째 리턴값 : 에러
func ValidateAdmin(jwtToken string) (bool, bool, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return false, false, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	accessToken := AccessToken{
		UserInfo: UserInfo{
			UserId:  claims["userId"].(int64),
			Email:   claims["email"].(string),
			IsAdmin: claims["isAdmin"].(bool),
			IsBlock: claims["isBlock"].(bool),
		},
	}

	if !ok {
		return false, false, errors.New("error parsing JWT claims")
	}

	return token.Valid, accessToken.UserInfo.IsAdmin, nil
}
