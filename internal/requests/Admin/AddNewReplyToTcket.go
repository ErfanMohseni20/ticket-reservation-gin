package admin

type AddNewReplyToTicket struct {
	Message string `json:"message" form:"message" binding:"required"`
}