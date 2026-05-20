package database

import (
	"fmt"
	"strings"
	"time"

	"github.com/ErfanMohseni20/ticket-reservation-gin/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect(cfg config.Config) error {
	var dsn string
	var db *gorm.DB
	var err error

	var logLevel logger.LogLevel
	if cfg.APPDEBUG == "true" {
		logLevel = logger.Info
	} else {
		logLevel = logger.Error
	}

	switch strings.ToLower(cfg.DBCONNECTION) {
	case "mysql", "mariadb":
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			cfg.DBUSERNAME, cfg.DBPASSWORD, cfg.DBHOST, cfg.DBPORT, cfg.DBDATABASE)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logLevel),
		})
	case "postgres", "postgresql", "pg", "":
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tehran",
			cfg.DBHOST, cfg.DBUSERNAME, cfg.DBPASSWORD, cfg.DBDATABASE, cfg.DBPORT)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logLevel),
		})
	default:
		// fallback to postgres
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tehran",
			cfg.DBHOST, cfg.DBUSERNAME, cfg.DBPASSWORD, cfg.DBDATABASE, cfg.DBPORT)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logLevel),
		})
	}

	if err != nil {
		return err
	}

	// set connection pool settings
	sqlDB, err := db.DB()
	if err == nil {
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetConnMaxLifetime(time.Hour)
	}

	DB = db
	return nil
}
