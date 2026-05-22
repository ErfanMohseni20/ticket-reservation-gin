package models

import "time"

type Ticket struct {
	ID        uint      `gorm:"primaryKey"`
	Title     string    `gorm:"type:varchar(255);not null;index"`
	CreatorID uint      `gorm:"not null;index"`
	Status    string    `gorm:"type:varchar(255);not null;default:'new'"`
	CreatedAt time.Time `gorm:"type:timestamptz;default:now()"`

	Creator       User          `gorm:"foreignKey:CreatorID;references:ID"`
	TicketReplies []TicketReply `gorm:"foreignKey:TicketID;references:ID"`
}

