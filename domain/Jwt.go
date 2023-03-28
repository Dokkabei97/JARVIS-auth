package domain

import "time"

type AccessToken struct {
	UserId  int64
	Email   string
	IsAdmin bool
	IsBlock bool
}

type RefreshToken struct {
	UserId int64
}

type Token struct {
	Id           int64
	userId       int64
	AccessToken  AccessToken
	RefreshToken RefreshToken
	CreatedAt    time.Time
}
