package domain

import "time"

type User struct {
	Id            int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	Email         string    `json:"email" gorm:"not null;unique"`
	EmailVerified time.Time `json:"emailVerified" gorm:"column:email_verified"`
	Name          string    `json:"name" gorm:"not null"`
	Password      string    `json:"password" gorm:"not null"`
	IsAdmin       bool      `json:"isAdmin" gorm:"not null;column:is_admin"`
	IsBlock       bool      `json:"isBlock" gorm:"not null;column:is_block"`
	LastLoginAt   time.Time `json:"lastLoginAt" gorm:"column:last_login_at"`
	CreatedAt     time.Time `json:"createdAt" gorm:"autoCreateTime;not null;column:created_at"`
	UpdatedAt     time.Time `json:"updatedAt" gorm:"autoUpdateTime;not null;column:updated_at"`
}
