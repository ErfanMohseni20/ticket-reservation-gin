package models

type City struct {
	ID uint `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(255);uniqueIndex;not null"`
	Terminals []Terminal `gorm:"foreignKey:CityId"`
}
