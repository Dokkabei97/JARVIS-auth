package domain

import (
	"time"
)

type JwtToken struct {
	Id           int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	UserId       int64     `json:"userId" gorm:"not null;column:user_id"`
	AccessToken  string    `json:"accessToken" gorm:"not null;column:access_token"`
	RefreshToken string    `json:"refreshToken" gorm:"not null;column:refresh_token"`
	CreatedAt    time.Time `json:"createdAt" gorm:"autoCreateTime;not null;column:created_at"`
}
