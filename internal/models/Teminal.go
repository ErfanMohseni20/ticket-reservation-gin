package models
type Terminal struct {
	ID      uint   `gorm:"primaryKey"`
	CityID  uint   `gorm:"not null;index"`
	Name    string `gorm:"type:varchar(255);not null"`
	City    City   `gorm:"foreignKey:CityID"`
}

func (Terminal) TableName() string {
	return "terminals"
}