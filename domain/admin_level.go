package domain

import "time"

type AdminLevel struct {
	Id        int64     `json:"id"gorm:"primary_key;auto_increment"`
	UserId    int64     `json:"userId"gorm:"not null;column:user_id"`
	Level     string    `json:"level"gorm:"not null"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime;not null;column:created_at"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime;not null;column:updated_at"`
}
