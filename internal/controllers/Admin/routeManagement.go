package admin

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ErfanMohseni20/ticket-reservation-gin/internal/database"
	"github.com/ErfanMohseni20/ticket-reservation-gin/internal/models"
	response "github.com/ErfanMohseni20/ticket-reservation-gin/internal/responses/Admin"
	request "github.com/ErfanMohseni20/ticket-reservation-gin/internal/requests/Admin"
	"github.com/gin-gonic/gin"
)

func RouteList(c *gin.Context) {
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

	var routes []models.Route
	result := database.DB.
		Preload("OriginTerminal.City").
		Preload("DestinationTerminal.City").
		Limit(perPage).
		Offset(offset).
		Find(&routes)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to fetch data from database"})
		return
	}

	var total int64
	database.DB.Model(&models.Route{}).Count(&total)

	var responseList []response.RouteResponse
	for _, route := range routes {
		responseList = append(responseList, response.RouteResponse{
			ID: route.ID,
			OriginTerminal: response.TerminalResponse{
				ID:   route.OriginTerminal.ID,
				Name: route.OriginTerminal.Name,
				City: response.CityResponse{
					ID:   route.OriginTerminal.City.ID,
					Name: route.OriginTerminal.City.Name,
				},
			},
			DestinationTerminal: response.TerminalResponse{
				ID:   route.DestinationTerminal.ID,
				Name: route.DestinationTerminal.Name,
				City: response.CityResponse{
					ID:   route.DestinationTerminal.City.ID,
					Name: route.DestinationTerminal.City.Name,
				},
			},
			Duration: route.Duration,
			Distance: route.Distance,
		})
	}

	totalPages := int((total + int64(perPage) - 1) / int64(perPage))

	c.JSON(http.StatusOK, gin.H{
		"message": "routes fetched successfully",
		"data":    responseList,
		"pagination": gin.H{
			"current_page": page,
			"per_page":     perPage,
			"total":        total,
			"total_pages":  totalPages,
		},
	})
}

func RouteCreate(c *gin.Context) {
	var req request.AddNewRouteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if req.OriginTerminalId == req.DestinationTerminalId {
		c.JSON(http.StatusBadRequest, gin.H{"message": "origin and destination terminals must be different"})
		return
	}

	var origin models.Terminal
	if err := database.DB.First(&origin, req.OriginTerminalId).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "selected origin terminal does not exist"})
		return
	}

	var dest models.Terminal
	if err := database.DB.First(&dest, req.DestinationTerminalId).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "selected destination terminal does not exist"})
		return
	}

	route := models.Route{
		OriginTerminalID:      req.OriginTerminalId,
		DestinationTerminalID: req.DestinationTerminalId,
		Duration:              req.Duration,
		Distance:              req.Distance,
	}

	if err := database.DB.Create(&route).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to create route"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "route created successfully"})
}

func RouteShow(c *gin.Context) {
	routeIDStr := c.Param("id")
	routeID, err := strconv.ParseUint(routeIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid route ID"})
		return
	}

	var route models.Route
	if err := database.DB.Preload("OriginTerminal.City").Preload("DestinationTerminal.City").First(&route, routeID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("route id %v not found", routeID)})
		return
	}

	resp := response.RouteResponse{
		ID: route.ID,
		OriginTerminal: response.TerminalResponse{
			ID:   route.OriginTerminal.ID,
			Name: route.OriginTerminal.Name,
			City: response.CityResponse{
				ID:   route.OriginTerminal.City.ID,
				Name: route.OriginTerminal.City.Name,
			},
		},
		DestinationTerminal: response.TerminalResponse{
			ID:   route.DestinationTerminal.ID,
			Name: route.DestinationTerminal.Name,
			City: response.CityResponse{
				ID:   route.DestinationTerminal.City.ID,
				Name: route.DestinationTerminal.City.Name,
			},
		},
		Duration: route.Duration,
		Distance: route.Distance,
	}

	c.JSON(http.StatusOK, gin.H{"message": "route fetched successfully", "data": resp})
}

func RouteUpdate(c *gin.Context) {
	routeIDStr := c.Param("id")
	routeID, err := strconv.ParseUint(routeIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid route ID"})
		return
	}

	var req request.UpdateRouteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var route models.Route
	if err := database.DB.First(&route, routeID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("route id %v not found", routeID)})
		return
	}

	if req.OriginTerminalId > 0 {
		var origin models.Terminal
		if err := database.DB.First(&origin, req.OriginTerminalId).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "selected origin terminal does not exist"})
			return
		}
		route.OriginTerminalID = req.OriginTerminalId
	}

	if req.DestinationTerminalId > 0 {
		var dest models.Terminal
		if err := database.DB.First(&dest, req.DestinationTerminalId).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "selected destination terminal does not exist"})
			return
		}
		route.DestinationTerminalID = req.DestinationTerminalId
	}

	if req.Duration != "" {
		route.Duration = req.Duration
	}
	if req.Distance > 0 {
		route.Distance = req.Distance
	}

	if err := database.DB.Save(&route).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to update route"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "route updated successfully"})
}

func RouteDelete(c *gin.Context) {
	routeIDStr := c.Param("id")
	routeID, err := strconv.ParseUint(routeIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid route ID"})
		return
	}

	var route models.Route
	if err := database.DB.First(&route, routeID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("route id %v not found", routeID)})
		return
	}

	if err := database.DB.Delete(&route).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to delete route"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "route deleted successfully"})
}
