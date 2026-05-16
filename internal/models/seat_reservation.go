package models

import "time"

type BusSeat struct {
	ID         uint   `gorm:"primaryKey"`
	BusID      uint   `gorm:"not null;index"`
	SeatNumber int    `gorm:"not null;check:seat_number > 0"`
	Status     string `gorm:"type:varchar(20);default:'available';check:status IN ('available','reserved','purchased','maintenance')"`
	Bus        Bus    `gorm:"foreignKey:BusID"`
}

type SeatReservation struct {
	ID          uint       `gorm:"primaryKey"`
	BusID       uint       `gorm:"not null;index"`
	BusSeatID   uint       `gorm:"not null;index"`
	UserID      uint       `gorm:"not null;index"`
	Status      string     `gorm:"type:varchar(255);not null;check:status IN ('reserved','purchased','canceled')"`
	ReservedAt  time.Time  `gorm:"type:timestamptz;default:now()"`
	PurchasedAt *time.Time `gorm:"type:timestamptz"`
	Bus         Bus        `gorm:"foreignKey:BusID"`
	BusSeat     BusSeat    `gorm:"foreignKey:BusSeatID"`
	User        User       `gorm:"foreignKey:UserID"`
}