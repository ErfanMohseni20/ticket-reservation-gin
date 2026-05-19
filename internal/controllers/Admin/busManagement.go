package admin

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/ErfanMohseni20/ticket-reservation-gin/internal/database"
	"github.com/ErfanMohseni20/ticket-reservation-gin/internal/models"
	request "github.com/ErfanMohseni20/ticket-reservation-gin/internal/requests/Admin"
	response "github.com/ErfanMohseni20/ticket-reservation-gin/internal/responses/Admin"
	"github.com/gin-gonic/gin"
)

func BusList(c *gin.Context) {
	perPageStr := c.DefaultQuery("per_page", "15")
	pageStr := c.DefaultQuery("page", "1")
	perPage, err := strconv.Atoi(perPageStr)
	if err != nil || perPage < 1 || perPage > 100 {
		perPage = 15
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}
	offset := (page - 1) * perPage
	var buses []models.Bus
	result := database.DB.Limit(perPage).Offset(offset).Find(&buses)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to fetched data from database"})
		return
	}
	var total int64
	database.DB.Model(&models.Bus{}).Count(&total)
	var responseList []response.BusResponse
	for _, bus := range buses {
		responseList = append(responseList, response.BusResponse{
			ID:               bus.ID,
			RouteID:          bus.RouteID,
			DepartureTime:    bus.DepartureTime.Format("2006-01-02 15:04:05"),
			ArrivalTime:      bus.ArrivalTime.Format("2006-01-02 15:04:05"),
			Capacity:         bus.Capacity,
			Price:            bus.Price,
			BusType:          bus.BusType,
			Corporation:      *bus.Corporation,
			SuperCorporation: *bus.SuperCorporation,
			ServiceNumber:    *bus.ServiceNumber,
			IsVIP:            bus.IsVIP,
		})
	}
	totalPages := int((total + int64(perPage) - 1) / int64(perPage))

	c.JSON(http.StatusOK, gin.H{
		"message": "buses fetched successfully",
		"data":    responseList,
		"pagination": gin.H{
			"current_page": page, "per_page": perPage,
			"total": total, "total_pages": totalPages,
		},
	})

}
func BusCreate(c *gin.Context) {
	var AddNewBusRequest request.AddNewBusRequest
	if err := c.ShouldBindJSON(&AddNewBusRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := database.DB.First(&models.Route{}, AddNewBusRequest.RouteID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("route id %v not exists", AddNewBusRequest.RouteID)})
		return
	}
	const layout = "2006-01-02 15:04"
	depTime, err := time.Parse(layout, AddNewBusRequest.DepartureTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid departure_time format, expected 'YYYY-MM-DD HH:MM:SS'"})
		return
	}
	arrTime, err := time.Parse(layout, AddNewBusRequest.ArrivalTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid arrival_time format, expected 'YYYY-MM-DD HH:MM:SS'"})
		return
	}

	bus := models.Bus{
		RouteID:          uint(AddNewBusRequest.RouteID),
		DepartureTime:    depTime,
		ArrivalTime:      arrTime,
		Capacity:         AddNewBusRequest.Capacity,
		Price:            AddNewBusRequest.Price,
		IsVIP:            AddNewBusRequest.IsVIP,
		Corporation:      &AddNewBusRequest.Corporation,
		SuperCorporation: &AddNewBusRequest.SuperCorporation,
		BusType:          AddNewBusRequest.BusType,
		ServiceNumber:    &AddNewBusRequest.ServiceNumber,
	}
	if err := database.DB.Create(&bus).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to create new bus "})
		return
	}
	for i := 1; i <= bus.Capacity; i++ {
		busSeat := models.BusSeat{
			BusID:      bus.ID,
			SeatNumber: i,
			Status:     "available",
		}
		if err := database.DB.Create(&busSeat).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to create seats for bus"})
			return
		}
	}
	c.JSON(http.StatusCreated, gin.H{"message": "bus created successfully"})
}
func BusShow(c *gin.Context) {
	busIdStr := c.Param("id")
	busid, err := strconv.ParseUint(busIdStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid bus ID"})
		return
	}
	var bus models.Bus
	if err := database.DB.First(&bus, busid).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("bus id %v not found", busid)})
		return
	}
	var responseFormat []response.BusResponse
	responseFormat = append(responseFormat, response.BusResponse{
		ID:               bus.ID,
		RouteID:          bus.RouteID,
		DepartureTime:    bus.DepartureTime.Format("2006-01-02 15:04:05"),
		ArrivalTime:      bus.ArrivalTime.Format("2006-01-02 15:04:05"),
		Capacity:         bus.Capacity,
		Price:            bus.Price,
		BusType:          bus.BusType,
		Corporation:      *bus.Corporation,
		SuperCorporation: *bus.SuperCorporation,
		ServiceNumber:    *bus.ServiceNumber,
		IsVIP:            bus.IsVIP,
	})
	c.JSON(http.StatusOK, gin.H{"message": "data fetched successfully", "data": responseFormat})
}
func BusUpdate(c *gin.Context) {
	busIdStr := c.Param("id")
	busId, err := strconv.ParseUint(busIdStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid bus id"})
	}
	var bus models.Bus
	if err := database.DB.First(&bus, busId).Error; err != nil {
		c.JSON(http.StatusFound, gin.H{"message": fmt.Sprintf("bus id %v not found", busId)})
		return
	}
	var UpdateBusRequest request.UpdateBusRequest
	if err := c.ShouldBindJSON(&UpdateBusRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if UpdateBusRequest.RouteID != nil {
		var route models.Route
		if err := database.DB.First(&route, *UpdateBusRequest.RouteID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "selected route does not exist"})
			return
		}
		bus.RouteID = uint(*UpdateBusRequest.RouteID)
	}
	const layout = "2006-01-02 15:04"
	if UpdateBusRequest.DepartureTime != "" {
		depTime, err := time.Parse(layout, UpdateBusRequest.DepartureTime)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid departure_time format, expected 'YYYY-MM-DD HH:MM:SS'"})
			return
		}
		bus.DepartureTime = depTime
	}
	if UpdateBusRequest.ArrivalTime != "" {
		arrTime, err := time.Parse(layout, UpdateBusRequest.ArrivalTime)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid arrival_time format, expected 'YYYY-MM-DD HH:MM:SS'"})
			return
		}
		bus.ArrivalTime = arrTime
	}
	if UpdateBusRequest.Price != 0 {
		bus.Price = UpdateBusRequest.Price
	}
	if UpdateBusRequest.Capacity != 0 {
		bus.Capacity = UpdateBusRequest.Capacity
	}
	if UpdateBusRequest.BusType != "" {
		bus.BusType = UpdateBusRequest.BusType
	}
	if UpdateBusRequest.Corporation != "" {
		bus.Corporation = &UpdateBusRequest.Corporation
	}
	if UpdateBusRequest.SuperCorporation != "" {
		bus.SuperCorporation = &UpdateBusRequest.SuperCorporation
	}
	bus.IsVIP = UpdateBusRequest.IsVIP
	if UpdateBusRequest.ServiceNumber != "" {
		bus.ServiceNumber = &UpdateBusRequest.ServiceNumber
	}
	if err := database.DB.Save(&bus).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to update bus"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "bus updated successfully"})

}
func BusDelete(c *gin.Context) {
	busIdStr := c.Param("id")
	busId, err := strconv.ParseUint(busIdStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid bus id"})
	}
	var bus models.Bus
	if err := database.DB.First(&bus, busId).Error; err != nil {
		c.JSON(http.StatusFound, gin.H{"message": fmt.Sprintf("bus id %v not found", busId)})
		return
	}
	if err := database.DB.Delete(&bus).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to delete bus from database"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "bus deleted successfully"})
}
func UpdateBusSeatStatus(c *gin.Context) {
	var UpdateBusSeatStatusRequest request.UpdateBusSeatStatusRequest
	if err := c.ShouldBindJSON(&UpdateBusSeatStatusRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"meesage": err.Error()})
		return
	}
	var bus models.Bus
	if err := database.DB.First(&bus, UpdateBusSeatStatusRequest.BusID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("bus id %v not found", UpdateBusSeatStatusRequest.BusID)})
		return
	}
	var busSeat models.BusSeat
	if err := database.DB.Model(&busSeat).Where("id = ? and bus_id = ?", UpdateBusSeatStatusRequest.SeatID, UpdateBusSeatStatusRequest.BusID).First(&busSeat).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("bus id %v haven't seat id %v", UpdateBusSeatStatusRequest.BusID, UpdateBusSeatStatusRequest.SeatID)})
		return
	}
	busSeat.Status=UpdateBusSeatStatusRequest.Status
	if err := database.DB.Save(&busSeat).Error;err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"message" : "failed to update bus seat status"})
		return
	}
	c.JSON(http.StatusOK,gin.H{"message" : "seat updated successfully"})
}
