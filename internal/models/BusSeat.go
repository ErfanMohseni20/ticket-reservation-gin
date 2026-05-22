package models

type BusSeatStatus string 

const (
	BusSeatAvailable  BusSeatStatus = "available"
	BusSeatReserved  BusSeatStatus = "reserved"
	BusSeatPurchased BusSeatStatus = "purchased"
	BusSeatMaintenance  BusSeatStatus = "maintenance"
)

type BusSeat struct {
	ID         uint   `gorm:"primaryKey"`
	BusID      uint   `gorm:"not null;index"`
	SeatNumber int    `gorm:"not null;check:seat_number > 0"`
	Status     string `gorm:"type:varchar(20);default:'available';check:status IN ('available','reserved','purchased','maintenance')"`
	Bus        Bus    `gorm:"foreignKey:BusID"`
}

