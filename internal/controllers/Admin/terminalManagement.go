package admin

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ErfanMohseni20/ticket-reservation-gin/internal/database"
	"github.com/ErfanMohseni20/ticket-reservation-gin/internal/models"
	request "github.com/ErfanMohseni20/ticket-reservation-gin/internal/requests/Admin"
	response "github.com/ErfanMohseni20/ticket-reservation-gin/internal/responses/Admin"
	"github.com/gin-gonic/gin"
)

func TerminalList(c *gin.Context) {
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

	var terminals []models.Terminal
	result := database.DB.
		Preload("City").
		Limit(perPage).
		Offset(offset).
		Find(&terminals)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to fetch data from database"})
		return
	}

	var total int64
	database.DB.Model(&models.Terminal{}).Count(&total)

	var responseList []response.TerminalResponse
	for _, terminal := range terminals {
		responseList = append(responseList, response.TerminalResponse{
			ID:   terminal.ID,
			Name: terminal.Name,
			City: response.CityResponse{
				ID:   terminal.City.ID,
				Name: terminal.City.Name,
			},
		})
	}

	totalPages := int((total + int64(perPage) - 1) / int64(perPage))

	c.JSON(http.StatusOK, gin.H{
		"message": "terminals fetched successfully",
		"data":    responseList,
		"pagination": gin.H{
			"current_page": page,
			"per_page":     perPage,
			"total":        total,
			"total_pages":  totalPages,
		},
	})
}

func TerminalCreate(c *gin.Context) {
	var req request.AddNewTerminalRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var city models.City
	if err := database.DB.First(&city, req.CityId).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "selected city does not exist"})
		return
	}

	terminal := models.Terminal{
		Name:   req.Name,
		CityID: city.ID,
	}

	if err := database.DB.Create(&terminal).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to create terminal"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "terminal created successfully"})
}

func TerminalUpdate(c *gin.Context) {
	terminalIDStr := c.Param("id")
	terminalID, err := strconv.ParseUint(terminalIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid terminal ID"})
		return
	}

	var req request.UpdateTerminalRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var terminal models.Terminal
	if err := database.DB.First(&terminal, terminalID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("terminal id %v not found", terminalID)})
		return
	}

	if req.Name != "" {
		terminal.Name = req.Name
	}

	if req.CityId > 0 {
		var city models.City
		if err := database.DB.First(&city, req.CityId).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "selected city does not exist"})
			return
		}
		terminal.CityID = city.ID
	}

	if err := database.DB.Save(&terminal).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to update terminal"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "terminal updated successfully"})
}

func TerminalDelete(c *gin.Context) {
	terminalIDStr := c.Param("id")
	terminalID, err := strconv.ParseUint(terminalIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid terminal ID"})
		return
	}

	var terminal models.Terminal
	if err := database.DB.First(&terminal, terminalID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("terminal id %v not found", terminalID)})
		return
	}

	if err := database.DB.Delete(&terminal).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to delete terminal"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "terminal deleted successfully"})
}