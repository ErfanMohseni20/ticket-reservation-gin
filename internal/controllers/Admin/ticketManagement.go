package admin

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/ErfanMohseni20/ticket-reservation-gin/internal/database"
	"github.com/ErfanMohseni20/ticket-reservation-gin/internal/helpers"
	"github.com/ErfanMohseni20/ticket-reservation-gin/internal/models"
	request "github.com/ErfanMohseni20/ticket-reservation-gin/internal/requests/Admin"
	response "github.com/ErfanMohseni20/ticket-reservation-gin/internal/responses/Admin"
	"github.com/gin-gonic/gin"
)

func TicketList(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	perPageStr := c.DefaultQuery("per_page", "15")

	page, _ := strconv.Atoi(pageStr)
	perPage, _ := strconv.Atoi(perPageStr)
	if page < 1 {
		page = 1
	}
	if perPage < 1 || perPage > 100 {
		perPage = 15
	}

	offset := (page - 1) * perPage
	var tickets []models.Ticket
	if err := database.DB.Preload("Creator").Limit(perPage).Offset(offset).Find(&tickets).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to fetched data from database"})
		return
	}
	var responseList []response.TicketResponse
	for _, ticket := range tickets {
		responseList = append(responseList, response.TicketResponse{
			ID:          ticket.ID,
			Title:       ticket.Title,
			CreatorName: ticket.Creator.FullName,
			Status:      ticket.Status,
			CreatedAt:   ticket.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	var total int64
	database.DB.Model(&models.Ticket{}).Count(&total)
	totalPages := int((total + int64(perPage) - 1) / int64(perPage))

	c.JSON(http.StatusOK, gin.H{
		"data": responseList,
		"pagination": gin.H{
			"current_page": page, "per_page": perPage,
			"total": total, "total_pages": totalPages,
		},
	})
}
func TicketClose(c *gin.Context) {
	ticketidstr := c.Param("id")
	ticketId, err := strconv.Atoi(ticketidstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid ticket id "})
	}
	var ticket models.Ticket
	if err := database.DB.First(&ticket, ticketId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("ticket id %v not found", ticketId)})
		return
	}
	ticket.Status = "close"
	if err := database.DB.Save(&ticket).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to update ticket status"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ticket updated successfully"})
}
func TicketShow(c *gin.Context) {
	ticketidstr := c.Param("id")
	ticketId, err := strconv.Atoi(ticketidstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid ticket id "})
	}
	var ticket models.Ticket
	if err := database.DB.Preload("TicketReplies").First(&ticket, ticketId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("ticket id %v not found", ticketId)})
		return
	}
	ticketResponse := response.TicketResponse{
		ID:          ticket.ID,
		Title:       ticket.Title,
		Status:      ticket.Status,
		CreatedAt:   ticket.CreatedAt.Format("2006-01-02 15:04:05"),
		CreatorName: ticket.Creator.Username,
	}
	var TicketRepliesResponse []response.TicketRepliesResponse
	for _, ticketreply := range ticket.TicketReplies {
		TicketRepliesResponse = append(TicketRepliesResponse, response.TicketRepliesResponse{
			ID:         ticketreply.ID,
			Message:    ticketreply.Message,
			SenderRole: ticketreply.SenderRole,
			CreatedAt:  ticketreply.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	data := map[string]interface{}{
		"ticket":         ticketResponse,
		"ticket_replies": TicketRepliesResponse,
	}
	c.JSON(http.StatusOK, gin.H{"message": "ticket fetched successfully", "data": data})
}
func TicketReply(c *gin.Context) {
	claims := helpers.MustGetUserFromContext(c)
	ticketidstr := c.Param("id")
	ticketId, err := strconv.Atoi(ticketidstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid ticket id "})
	}
	var AddNewReplyToTicket request.AddNewReplyToTicket
	if err := c.ShouldBindJSON(&AddNewReplyToTicket);err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"message" :err.Error()})
		return
	}
	var ticket models.Ticket
	if err := database.DB.First(&ticket, ticketId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("ticket id %v not found", ticketId)})
		return
	}
	ticketReply := models.TicketReply{
		TicketID:   ticket.ID,
		Message:    AddNewReplyToTicket.Message,
		SenderID:   &claims.UserID,
		SenderRole: "customer",
		CreatedAt:  time.Now(),
	}
	if err := database.DB.Create(&ticketReply).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to create ticket reply"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "ticket reply created successfully"})
}
