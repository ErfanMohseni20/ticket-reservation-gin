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
	"golang.org/x/crypto/bcrypt"
)

func UsersList(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	perPageStr := c.DefaultQuery("per_page", "15")

	page, _ := strconv.Atoi(pageStr)
	perPage, _ := strconv.Atoi(perPageStr)
	if page < 1 { page = 1 }
	if perPage < 1 || perPage > 100 { perPage = 15 }

	offset := (page - 1) * perPage

	var users []models.User
	result := database.DB.Limit(perPage).Offset(offset).Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to fetch users"})
		return
	}

	var total int64
	database.DB.Model(&models.User{}).Count(&total)

	var responseList []response.UsersResponse
	for _, user := range users {
		responseList = append(responseList, response.UsersResponse{
			ID: user.ID, Username: user.Username,
			Fullname: user.FullName,
			CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	totalPages := int((total + int64(perPage) - 1) / int64(perPage))

	c.JSON(http.StatusOK, gin.H{
		"data": responseList,
		"pagination": gin.H{
			"current_page": page, "per_page": perPage,
			"total": total, "total_pages": totalPages,
		},
	})
}
func UserShow(c *gin.Context) {
	userid := c.Param("id")
	var User models.User
	if err := database.DB.Model(&User).Where("id = ?", userid).First(&User).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("user %v not found", userid)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": response.UsersResponse{
		ID : User.ID,
		Fullname: User.FullName,
		Username: User.Username,
		CreatedAt: User.CreatedAt.Format("2006-01-02 15:04:05"),
	}})

}
func UserCreate(c *gin.Context) {
	var UserCreateRequest request.AddNewUserReqeust
	if err := c.ShouldBindJSON(&UserCreateRequest);err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"message" : err.Error()})
		return
	}
	hashedPassword,err := bcrypt.GenerateFromPassword([]byte(UserCreateRequest.Password),bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"message" : "failed to generate hashed password"})
		return
	}
	user := models.User{
		FullName : UserCreateRequest.FullName,
		Username : UserCreateRequest.UserName,
		HashedPassword: string(hashedPassword),
		CreatedAt : time.Now(),
	}
	if err:=database.DB.Create(&user).Error;err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"message" : "failed to create user"})
		return
	}
	c.JSON(http.StatusCreated,gin.H{"message" : "user created successfully"})
}
func UserUpdate(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid user ID"})
		return
	}

	var req request.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	if req.FullName != "" {
		user.FullName = req.FullName
	}
	if req.UserName != "" {
		user.Username = req.UserName
	}
	if req.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to hash password"})
			return
		}
		user.HashedPassword = string(hashedPassword)
		user.PasswordChangedAt = time.Now()
	}

	if err := database.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user updated successfully",
		"data": response.UsersResponse{
			ID:        user.ID,
			Username:  user.Username,
			Fullname:  user.FullName,
			CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}
func UserDelete(c *gin.Context) {
	userid := c.Param("id")
	var User models.User 
	if err := database.DB.Model(&User).Where("id = ?",userid).First(&User);err != nil {
		c.JSON(http.StatusNotFound,gin.H{"message" : fmt.Sprintf("user %v not found",userid)})
		return
	}
	if err := database.DB.Model(&User).Where("id = ?",userid).Delete(&User);err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"message" : "failed to delete user from database"})
		return
	}
	c.JSON(http.StatusOK,gin.H{"message" : "user deleted successfully"})
}
