package models

import "time"

type Route struct {
	ID                     uint      `gorm:"primaryKey"`
	OriginTerminalID       uint      `gorm:"not null;index"`
	DestinationTerminalID  uint      `gorm:"not null;index"`
	Duration               string    `gorm:"type:interval;not null"`
	Distance               int       `gorm:"not null;check:distance > 0"`
	OriginTerminal         Terminal  `gorm:"foreignKey:OriginTerminalID"`
	DestinationTerminal    Terminal  `gorm:"foreignKey:DestinationTerminalID"`
}

type Bus struct {
	ID                uint      `gorm:"primaryKey"`
	RouteID           uint      `gorm:"not null;index"`
	DepartureTime     time.Time `gorm:"type:timestamptz;not null"`
	ArrivalTime       time.Time `gorm:"type:timestamptz;not null"`
	Capacity          int       `gorm:"not null;check:capacity > 0"`
	Price             int       `gorm:"not null;check:price > 0"`
	BusType           string    `gorm:"type:varchar(255);not null"`
	Corporation       *string   `gorm:"type:varchar(255)"`
	SuperCorporation  *string   `gorm:"type:varchar(255)"`
	ServiceNumber     *string   `gorm:"type:varchar(255)"`
	IsVIP             bool      `gorm:"default:false"`
	Route             Route     `gorm:"foreignKey:RouteID"`
}