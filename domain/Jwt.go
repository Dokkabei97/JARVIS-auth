package domain

import "time"

type AccessToken struct {
	UserId  int64  `json:"userId"`
	Email   string `json:"email"`
	IsAdmin bool   `json:"isAdmin"`
	IsBlock bool   `json:"isBlock"`
}

type RefreshToken struct {
	UserId int64 `json:"userId"`
}

type Token struct {
	Id           int64        `json:"id"`
	UserId       int64        `json:"userId"`
	AccessToken  AccessToken  `json:"accessToken"`
	RefreshToken RefreshToken `json:"refreshToken"`
	CreatedAt    time.Time    `json:"createdAt"`
}
