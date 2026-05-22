package customer

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ErfanMohseni20/ticket-reservation-gin/internal/database"
	"github.com/ErfanMohseni20/ticket-reservation-gin/internal/helpers"
	"github.com/ErfanMohseni20/ticket-reservation-gin/internal/models"
	request "github.com/ErfanMohseni20/ticket-reservation-gin/internal/requests/Customer"
	response "github.com/ErfanMohseni20/ticket-reservation-gin/internal/responses/Customer"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func ReserveSeat(c *gin.Context) {
	claims := helpers.MustGetUserFromContext(c)

	var req request.ReserveSeat
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var reservation *models.SeatReservation
	var seat *models.BusSeat

	err := database.DB.Transaction(func(tx *gorm.DB) error {
		var bus models.Bus
		if err := tx.Preload("Seats").First(&bus, req.BusID).Error; err != nil {
			return fmt.Errorf("bus not found")
		}

		var availableSeat models.BusSeat
		if err := tx.
			Where("bus_id = ? AND status = ?", req.BusID, "available").
			Order("seat_number ASC").
			Clauses(clause.Locking{Strength: "UPDATE", Options: "SKIP LOCKED"}).
			First(&availableSeat).Error; err != nil {
			return fmt.Errorf("no available seats found")
		}

		newReservation := models.SeatReservation{
			BusID:     bus.ID,
			BusSeatID: availableSeat.ID,
			UserID:    claims.UserID,
			Status:    "reserved",
			ReservedAt: time.Now(),
		}
		if err := tx.Create(&newReservation).Error; err != nil {
			return err
		}

		if err := tx.Model(&availableSeat).Update("status", "reserved").Error; err != nil {
			return err
		}

		reservation = &newReservation
		seat = &availableSeat
		return nil
	})

	if err != nil {
		if err.Error() == "bus not found" {
			c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("bus id %v not found", req.BusID)})
			return
		}
		if err.Error() == "no available seats found" {
			c.JSON(http.StatusConflict, gin.H{"message": "no available seats on this bus"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "reservation failed: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "seat reserved seccessfully",
		"data": response.ReservationResponse{
			ID:         reservation.ID,
			BusID:      reservation.BusID,
			SeatNumber: seat.SeatNumber,
			Status:     reservation.Status,
			ReservedAt: reservation.ReservedAt.Format("2006-01-02 15:04:05"),
		}})
}
func MyReserveList(c *gin.Context) {
	claims := helpers.MustGetUserFromContext(c)
	var MyReserveList []models.SeatReservation
	if err := database.DB.Preload("User").Preload("BusSeat").Model(&MyReserveList).Where("user_id = ?", claims.UserID).Find(&MyReserveList).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "there is no record for you"})
		return
	}
	var responseList []response.ReservationResponse
	for _, reservation := range MyReserveList {
		responseList = append(responseList, response.ReservationResponse{
			ID:         reservation.ID,
			BusID:      reservation.BusID,
			SeatNumber: reservation.BusSeat.SeatNumber,
			Status:     reservation.Status,
			ReservedAt: reservation.ReservedAt.Format("2006-01-02 15:04:05"),
		})
	}
	c.JSON(http.StatusOK, gin.H{"message": "data fetched successfully", "data": responseList})
}
func ChnageStatus(c *gin.Context) {
	claims := helpers.MustGetUserFromContext(c)
	var ReserveChangeStatus request.ReserveChangeStatus
	if err := c.ShouldBindJSON(&ReserveChangeStatus); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	var Reserve models.SeatReservation
	if err := database.DB.Model(&Reserve).Where("id = ? and user_id = ? and status = ?", ReserveChangeStatus.ReserveID, claims.UserID, "reserved").First(&Reserve).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("reserve id %v not found or not blong to you or something else ", ReserveChangeStatus.ReserveID)})
		return
	}
	if ReserveChangeStatus.Status == "purchased" {
		Reserve.Status = "purchased"
		Reserve.PurchasedAt = time.Now()
	}
	if (ReserveChangeStatus.Status == "canceled" ){
		Reserve.Status = "canceled"
	}
	if err := database.DB.Save(&Reserve).Error;err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"message" : "failed to update reserve record"})
		return
	}
	c.JSON(http.StatusOK,gin.H{"message" : "record updated successfully"})
}
