package main

import (
	"log"

	"github.com/ErfanMohseni20/ticket-reservation-gin/internal/config"
	"github.com/ErfanMohseni20/ticket-reservation-gin/internal/database"
	"github.com/ErfanMohseni20/ticket-reservation-gin/internal/database/migrations"
	"github.com/ErfanMohseni20/ticket-reservation-gin/internal/helpers"
	"github.com/ErfanMohseni20/ticket-reservation-gin/internal/models"
	"github.com/ErfanMohseni20/ticket-reservation-gin/internal/routers"
	"go.uber.org/zap"
)

func main(){
	// load config
	config,err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("failed to load config;",err)
	} 
	//----------------------------------------------
	// setup logger 
	logger , _ := zap.NewProduction()
	defer logger.Sync()
	logger.Info("starting server",zap.String("app", config.APPNAME))

	//---------------------------------------------------------
	//connect to database and run migrations 
	if err := database.Connect(config); err != nil {
		logger.Fatal("DB Connection failed",zap.Error(err))
	}
	if database.DB != nil {
		if err := migrations.AddSenderIdToTicketReplies(database.DB);err != nil {
			logger.Fatal("failed to run migration",zap.Error(err))
		}
		if err := database.DB.AutoMigrate(&models.User{},&models.Bus{},&models.BusSeat{},&models.City{},&models.Route{},&models.SeatReservation{},&models.Terminal{},&models.Ticket{},&models.TicketReply{}); err != nil {
			logger.Fatal("auto migrate failed",zap.Error(err))
		}
		logger.Info("migrations completed")
	} else {
		logger.Info("DB connection established but no SQL DB detected; skipping SQL migrations")
	}
	//----------------------------------------
	helpers.JWTSecret = []byte(config.JWTTOKEN)
	if len(helpers.JWTSecret) < 32 {
		log.Fatal("JWT_SECRET must be at least 32 characters for HS256")
	}
	//----------------------------------------------
	// setup router and start server
	router := routers.RoutersSetup()
	portAddress := ":"+config.APPPORT
	logger.Info("server running",zap.String("address",portAddress))
	if err := router.Run(portAddress);err != nil {
		logger.Fatal("server failed",zap.Error(err))
	}
}
 