package models


type Route struct {
	ID                    uint     `gorm:"primaryKey"`
	OriginTerminalID      uint     `gorm:"not null;index"`
	DestinationTerminalID uint     `gorm:"not null;index"`
	Duration              string   `gorm:"type:interval;not null"`
	Distance              int      `gorm:"not null;check:distance > 0"`
	OriginTerminal        Terminal `gorm:"foreignKey:OriginTerminalID"`
	DestinationTerminal   Terminal `gorm:"foreignKey:DestinationTerminalID"`
}

