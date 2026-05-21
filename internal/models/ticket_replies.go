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

type TicketReply struct {
	ID         uint      `gorm:"primaryKey"`
	TicketID   uint      `gorm:"not null;index"`
	Message    string    `gorm:"type:text"`
	SenderRole string    `gorm:"type:varchar(50);not null;column:sender_role"`
	SenderID   *uint     `gorm:"index" json:"sender_id"`
	CreatedAt  time.Time `gorm:"type:timestamptz;default:now()"`

	Sender *User  `gorm:"foreignKey:SenderID" json:"-"`
	Ticket Ticket `gorm:"foreignKey:TicketID;references:ID"`
}
