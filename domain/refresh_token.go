package domain

import (
	"github.com/golang-jwt/jwt/v5"
)

type RefreshToken struct {
	UserId int64 `json:"userId"`
	jwt.RegisteredClaims
}
