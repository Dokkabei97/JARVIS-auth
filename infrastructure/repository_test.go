package infrastructure

import (
	"github.com/go-playground/assert/v2"
	"github.com/golang-jwt/jwt/v5"
	"is-deploy-auth/domain"
	"testing"
	"time"
)

var secretKey = []byte("CuZXj3gUuw1ljge86FySVrDYtx2wCqa9wovAevj26UE=")

func TestMySqlRepository_Save(t *testing.T) {
	krLocation, _ := time.LoadLocation("Asia/Seoul")
	krTime := time.Date(2023, 3, 29, 13, 15, 0, 0, krLocation)
	expiresAt := jwt.NewNumericDate(krTime)

	userInfo := domain.UserInfo{
		UserId:  1,
		Email:   "zxc",
		IsAdmin: false,
		IsBlock: false,
	}
	accessToken := domain.AccessToken{
		UserInfo: userInfo,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: expiresAt,
			Issuer:    "test",
		},
	}
	access := jwt.NewWithClaims(jwt.SigningMethodHS256, accessToken)
	refreshToken := domain.RefreshToken{
		UserId: 1,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: expiresAt,
			Issuer:    "test",
		},
	}
	refresh := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshToken)

	at, _ := access.SignedString(secretKey)
	//fmt.Printf("accessToken => %s\n", at)

	rt, _ := refresh.SignedString(secretKey)
	//fmt.Printf("refreshToken => %s\n", rt)

	token := domain.Token{
		UserId:       1,
		AccessToken:  at,
		RefreshToken: rt,
	}

	assert.Equal(t, token.AccessToken, at)
	assert.Equal(t, token.RefreshToken, rt)
}
