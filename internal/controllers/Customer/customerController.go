package customer

import (
	"net/http"
	"github.com/ErfanMohseni20/ticket-reservation-gin/internal/database"
	"github.com/ErfanMohseni20/ticket-reservation-gin/internal/helpers"
	"github.com/ErfanMohseni20/ticket-reservation-gin/internal/models"
	response "github.com/ErfanMohseni20/ticket-reservation-gin/internal/responses/Customer"
	"github.com/gin-gonic/gin"
)

func UpdateProfile(c *gin.Context) {
	claims := helpers.MustGetUserFromContext(c)
	if fullName := c.PostForm("full_name"); fullName != "" {
		if err := helpers.ValidateFullName(fullName); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	updates := make(map[string]interface{})
	if fullName := c.PostForm("full_name"); fullName != "" {
		updates["full_name"] = fullName
	}

	if len(updates) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no fields to update"})
		return
	}

	if err := database.DB.Model(&models.User{}).Where("id = ?", claims.UserID).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update profile"})
		return
	}

	var updatedUser models.User
	if err := database.DB.First(&updatedUser, claims.UserID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch updated user"})
		return
	}

	c.JSON(http.StatusOK, response.ProfileResponse{
		Fullname:  updatedUser.FullName,
		Username:  updatedUser.Username,
		CreatedAt: updatedUser.CreatedAt.Format("2006-01-02 15:04:05"),
	})
}

func Profile(c *gin.Context) {
	claims := helpers.MustGetUserFromContext(c)

	var user models.User
	if err := database.DB.First(&user, claims.UserID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch user profile"})
		return
	}

	c.JSON(http.StatusOK, response.ProfileResponse{
		Fullname:  user.FullName,
		Username:  user.Username,
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
	})
}
