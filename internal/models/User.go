package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID                uint           `gorm:"primaryKey;autoIncrement"`
	Username          string         `gorm:"type:varchar(255);uniqueIndex;not null"`
	HashedPassword    string         `gorm:"type:varchar(255);uniqueIndex;not null"`
	FullName          string         `gorm:"type:varchar(255);not null"`
	AvatarURL         string         `gorm:"type:varchar(255)"`
	PasswordChangedAt time.Time      `gorm:"type:timestamptz;default:'0001-01-01 00:00:00Z'"`
	CreatedAt         time.Time      `gorm:"type:timestamptz;default:now()"`
	DeletedAt         gorm.DeletedAt `gorm:"index"`
}
