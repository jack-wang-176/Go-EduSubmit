package model

import (
	"time"

	"gorm.io/gorm"
)

type RefreshToken struct {
	gorm.Model

	UserID   uint   `gorm:"not null;index" json:"user_id"`
	UserName string `gorm:"not null" json:"user_name"`

	Token     string    `gorm:"type:varchar(512);not null;uniqueIndex" json:"token"`
	ExpiresAt time.Time `gorm:"not null" json:"expires_at"`

	Revoked bool `gorm:"default:false;not null" json:"revoked"`
}
