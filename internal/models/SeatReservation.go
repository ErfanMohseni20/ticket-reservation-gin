package models

import "time"
type SeatReservationStatus string

const (
	SeatReserved  SeatReservationStatus = "reserved"
	SeatPurchased SeatReservationStatus = "purchased"
	SeatCanceled  SeatReservationStatus = "canceled"
)

type SeatReservation struct {
	ID          uint      `gorm:"primaryKey"`
	BusID       uint      `gorm:"not null;index"`
	BusSeatID   uint      `gorm:"not null;index"`
	UserID      uint      `gorm:"not null;index"`
	Status      SeatReservationStatus    `gorm:"type:varchar(255);not null;check:status IN ('reserved','purchased','canceled')"`
	ReservedAt  time.Time `gorm:"type:timestamptz;default:now()"`
	PurchasedAt time.Time `gorm:"type:timestamptz"`
	Bus         Bus       `gorm:"foreignKey:BusID"`
	BusSeat     BusSeat   `gorm:"foreignKey:BusSeatID"`
	User        User      `gorm:"foreignKey:UserID"`
}
