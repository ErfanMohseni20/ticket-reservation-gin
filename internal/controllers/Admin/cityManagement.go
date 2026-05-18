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

func CityList(c *gin.Context) {
	perPageStr := c.DefaultQuery("per_page", "15")
	pageStr := c.DefaultQuery("page", "1")
	perPage, _ := strconv.Atoi(perPageStr)
	page, _ := strconv.Atoi(pageStr)
	if page < 1 {
		page = 1
	}
	if perPage < 1 || perPage > 100 {
		perPage = 15
	}
	offset := (page - 1) * perPage
	var cities []models.City
	result := database.DB.Limit(perPage).Offset(offset).Find(&cities)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to fetch data from database"})
		return
	}
	var total int64
	database.DB.Model(&models.City{}).Count(&total)
	var ResponseList []response.CityResponse
	for _, city := range cities {
		ResponseList = append(ResponseList, response.CityResponse{
			ID:   city.ID,
			Name: city.Name,
		})
	}
	totalPages := int((total + int64(perPage) - 1) / int64(perPage))
	c.JSON(http.StatusOK, gin.H{
		"message": "cities fetched successfully from database",
		"data":    ResponseList,
		"paginitaion": gin.H{
			"current_page": page,
			"per_page":     perPage,
			"total":        total,
			"total_pages":  totalPages,
		}})
}
func CityCreate(c *gin.Context) {
	var AddNewCityRequest request.AddNewCityRequest
	if err := c.ShouldBindJSON(&AddNewCityRequest);err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"message" : err.Error()})
		return
	}
	city := models.City{
		Name: AddNewCityRequest.Name,
	}
	if err := database.DB.Create(&city).Error;err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"message" : "failed to create new city"})
		return
	}
	c.JSON(http.StatusCreated,gin.H{"message" : "city created successfully"})
}
func CityUpdate(c *gin.Context) {
	cityid := c.Param("id")
	cityID, _ := strconv.ParseUint(cityid, 10, 32)
	var UpdateCityRequest request.UpdateCityRequest
	var city models.City
	if err := c.ShouldBindJSON(&UpdateCityRequest);err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"message" : err.Error()})
		return
	}
	if err := database.DB.Model(&city).Where("id = ?",cityID).First(&city);err != nil {
		c.JSON(http.StatusNotFound,gin.H{"message" : fmt.Sprintf("city id %v not found",cityID)})
		return
	}
	if UpdateCityRequest.Name != "" {
		city.Name = UpdateCityRequest.Name
	}

	if err := database.DB.Save(&city).Error;err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"message" : "failed to update city"})
		return
	}
	c.JSON(http.StatusOK,gin.H{"message" : "city updated successfully"})
}
func CityDelete(c *gin.Context) {
	cityid := c.Param("id")
	var city models.City
	if err := database.DB.Model(&city).Where("id = ?",cityid).First(&city);err != nil {
		c.JSON(http.StatusNotFound,gin.H{"message" : fmt.Sprintf("city id %v not found",cityid)})
		return
	}
	if err := database.DB.Model(&city).Where("id = ?",cityid).Delete(&city);err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"message" : fmt.Sprintf("failed to delete city %v from ",cityid)})
		return
	}
	c.JSON(http.StatusOK,gin.H{"message" : "city delete successfully"})
}
