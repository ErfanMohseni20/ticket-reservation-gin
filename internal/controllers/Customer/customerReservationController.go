package customer

import (
	"fmt"
	"net/http"
	"strconv"
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
			BusID:      bus.ID,
			BusSeatID:  availableSeat.ID,
			UserID:     claims.UserID,
			Status:     models.SeatReserved,
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
			Status:     string(reservation.Status),
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
			Status:     string(reservation.Status),
			ReservedAt: reservation.ReservedAt.Format("2006-01-02 15:04:05"),
		})
	}
	c.JSON(http.StatusOK, gin.H{"message": "data fetched successfully", "data": responseList})
}
func ChangeStatus(c *gin.Context) {
	claims := helpers.MustGetUserFromContext(c)
	var req request.ReserveChangeStatus
	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	allowed := map[models.SeatReservationStatus]bool{
		models.SeatPurchased: true,
		models.SeatCanceled:  true,
	}
	if !allowed[models.SeatReservationStatus(req.Status)] {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid status"})
		return
	}

	var reserve models.SeatReservation
	if err := database.DB.
		Where("id = ? AND user_id = ? AND status = ?", 
			req.ReserveID, 
			claims.UserID, 
			models.SeatReserved).
		First(&reserve).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "reservation not found"})
		return
	}

	reserve.Status = models.SeatReservationStatus(req.Status)
	
	if reserve.Status == models.SeatPurchased {
		reserve.PurchasedAt = time.Now() 
	}

	if err := database.DB.Save(&reserve).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to update"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "status updated successfully"})
}
func History(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	perPageStr := c.DefaultQuery("per_page", "15")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}
	perPage, err := strconv.Atoi(perPageStr)
	if err != nil || perPage < 1 || perPage > 100 {
		perPage = 15
	}
	offset := (page - 1) * perPage
	cliams := helpers.MustGetUserFromContext(c)
	var reservations []models.SeatReservation
	if err := database.DB.Preload("Bus.Route.OriginTerminal").
		Preload("Bus.Route.DestinationTerminal").Model(&reservations).Limit(perPage).Offset(offset).Where("user_id = ? and status = ?", cliams.UserID, models.SeatReserved).Find(&reservations).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "there aren't any record for you"})
		return
	}
	var total int64
	database.DB.Model(&models.SeatReservation{}).Count(&total)
	var responseList []response.ReservationHistory
	for _, counter := range reservations {
		responseList = append(responseList, response.ReservationHistory{
			ID:                      counter.ID,
			OriginTerminalName:      counter.Bus.Route.OriginTerminal.Name,
			DestinationTerminalName: counter.Bus.Route.DestinationTerminal.Name,
			Status:                  string(counter.Status),
			ReservedAt:              counter.ReservedAt.Format("2006-01-02 15:04:05"),
			PurchasedAt:             counter.PurchasedAt.Format("2006-01-02 15:04:05"),
		})
	}
	totalPages := int((total + int64(perPage) - 1) / int64(perPage))

	c.JSON(http.StatusOK, gin.H{
		"data": responseList,
		"pagination": gin.H{
			"current_page": page,
			"per_page":     perPage,
			"total":        total,
			"total_pages":  totalPages,
		},
	})
}
