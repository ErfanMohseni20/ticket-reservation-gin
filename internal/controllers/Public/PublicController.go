package public

import (
	"net/http"
	"strconv"

	"github.com/ErfanMohseni20/ticket-reservation-gin/internal/database"
	"github.com/ErfanMohseni20/ticket-reservation-gin/internal/models"
	busresponse "github.com/ErfanMohseni20/ticket-reservation-gin/internal/responses/Admin"
	response "github.com/ErfanMohseni20/ticket-reservation-gin/internal/responses/Public"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func BusList(c *gin.Context) {
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

	var buses []models.Bus
	if err := database.DB.
		Preload("Route.OriginTerminal").
		Preload("Route.DestinationTerminal").
		Preload("Seats", func(db *gorm.DB) *gorm.DB {
			return db.Where("status = ?", "available")
		}).
		Limit(perPage).
		Offset(offset).
		Find(&buses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to fetch buses"})
		return
	}

	var total int64
	database.DB.Model(&models.Bus{}).Count(&total)

	var responseList []response.BusResponse
	for _, bus := range buses {
		responseList = append(responseList, response.BusResponse{
			ID:                      bus.ID,
			OriginTerminalName:      bus.Route.OriginTerminal.Name,
			DestinationTerminalName: bus.Route.DestinationTerminal.Name,
			AvailableSeatsCount:     int64(len(bus.Seats)),
			BusData: busresponse.BusResponse{
				RouteID:          bus.RouteID,
				DepartureTime:    bus.DepartureTime.Format("2006-01-02 15:04"),
				ArrivalTime:      bus.ArrivalTime.Format("2006-01-02 15:04"),
				Capacity:         bus.Capacity,
				Price:            bus.Price,
				BusType:          bus.BusType,
				Corporation:      getStringValue(bus.Corporation),
				SuperCorporation: getStringValue(bus.SuperCorporation),
				ServiceNumber:    getStringValue(bus.ServiceNumber),
				IsVIP:            bus.IsVIP,
			},
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

func getStringValue(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
