package database

import (
	"fmt"

	"github.com/ErfanMohseni20/ticket-reservation-gin/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect(config config.Config) error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tehran",
		config.DBHOST, config.DBUSERNAME, config.DBPASSWORD, config.DBDATABASE, config.DBPORT)
	var logLevel logger.LogLevel
	if config.APPDEBUG == "true" {
		logLevel = logger.Info
	} else {
		logLevel = logger.Error
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})
	if err != nil {
		return err
	}
	DB = db
	return nil
}
