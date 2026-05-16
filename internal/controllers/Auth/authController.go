package Auth

import (
	"net/http"
	"time"

	"github.com/ErfanMohseni20/ticket-reservation-gin/internal/database"
	"github.com/ErfanMohseni20/ticket-reservation-gin/internal/helpers"
	"github.com/ErfanMohseni20/ticket-reservation-gin/internal/models"
	request "github.com/ErfanMohseni20/ticket-reservation-gin/internal/requests/Auth"
	response "github.com/ErfanMohseni20/ticket-reservation-gin/internal/responses/Auth"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var LoginRequest request.LoginRequest
	if err := c.ShouldBindJSON(&LoginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user models.User
	if err := database.DB.Where("username = ?", LoginRequest.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credetianls"})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(LoginRequest.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}
	accessToken, err := helpers.GenerateToken(user.ID, user.Username, 24*time.Hour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "token generation failed"})
		return
	}
	refreshToken, err := helpers.GenerateRefreshToken(user.ID, user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "refresh token generation failed"})
		return
	}
	c.JSON(http.StatusOK, response.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Message:      "login successfully",
		User: struct {
			ID       uint   `json:"id"`
			Username string `json:"username"`
			FullName string `json:"full_name"`
		}{
			ID:       user.ID,
			Username: user.Username,
			FullName: user.FullName,
		},
	})
}
func Register(c *gin.Context) {
	var RegisterRequest request.RegisterRequest
	if err := c.ShouldBindJSON(&RegisterRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(RegisterRequest.Password), bcrypt.DefaultCost)
	user := models.User{
		Username:       RegisterRequest.Username,
		HashedPassword: string(hashedPassword),
		FullName:       RegisterRequest.Fullname,
		CreatedAt: time.Now(),
		AvatarURL: "/avatars/default.jpg",

	}
	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	accessToken, err := helpers.GenerateToken(user.ID, user.Username, 24*time.Hour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "token generation failed"})
		return
	}
	refreshToken, err := helpers.GenerateRefreshToken(user.ID, user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "token generation failed"})
		return
	}
	c.JSON(http.StatusCreated, response.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Message: "register successfully",
		User: struct {
			ID       uint   "json:\"id\""
			Username string "json:\"username\""
			FullName string "json:\"full_name\""
		}{
			ID : user.ID,
			Username: user.Username,
			FullName: user.FullName,
		},
	})
}
