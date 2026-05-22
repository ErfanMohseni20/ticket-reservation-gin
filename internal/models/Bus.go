package models

import "time"

type Bus struct {
	ID               uint      `gorm:"primaryKey"`
	RouteID          uint      `gorm:"not null;index"`
	DepartureTime    time.Time `gorm:"type:timestamptz;not null"`
	ArrivalTime      time.Time `gorm:"type:timestamptz;not null"`
	Capacity         int       `gorm:"not null;check:capacity > 0"`
	Price            int       `gorm:"not null;check:price > 0"`
	BusType          string    `gorm:"type:varchar(255);not null"`
	Corporation      *string   `gorm:"type:varchar(255)"`
	SuperCorporation *string   `gorm:"type:varchar(255)"`
	ServiceNumber    *string   `gorm:"type:varchar(255)"`
	Route            Route     `gorm:"foreignKey:RouteID"`
	IsVIP            bool      `gorm:"column:is_vip;default:false"`
	Seats            []BusSeat `gorm:"foreignKey:BusID"`
}
