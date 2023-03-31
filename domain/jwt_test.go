package domain

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var testSecretKey = []byte("CuZXj3gUuw1ljge86FySVrDYtx2wCqa9wovAevj26UE=")

func TestGenerateAccessToken(t *testing.T) {
	krTime := time.Now().Add(time.Hour).In(time.FixedZone("KST", 9*60*60))
	expiresAt := jwt.NewNumericDate(krTime)

	userInfo := UserInfo{
		UserId:  1,
		Email:   "zxc",
		IsAdmin: false,
		IsBlock: false,
	}
	accessToken := AccessToken{
		UserInfo: userInfo,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: expiresAt,
			Issuer:    "test",
		},
	}
	access := jwt.NewWithClaims(jwt.SigningMethodHS256, accessToken)
	at, _ := access.SignedString(testSecretKey)

	assert.NotEmpty(t, at)
}

func TestValidateAccessToken(t *testing.T) {
	const OLD_TOKEN = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySW5mbyI6eyJ1c2VySWQiOjEsImVtYWlsIjoienhjIiwiaXNBZG1pbiI6ZmFsc2UsImlzQmxvY2siOmZhbHNlfSwiaXNzIjoidGVzdCIsImV4cCI6MTY4MDA2MzMwMH0.XzscMGiPKD-QcVkmgo2JMrElKnoknXDT6icBogYGEhA\n"

	token, _ := jwt.Parse(OLD_TOKEN, func(token *jwt.Token) (interface{}, error) {
		return testSecretKey, nil
	})

	assert.False(t, token.Valid)
}
