package customer

type AddnewTicketRequest struct {
	Message string `json:"message" form:"message" binding:"required"`
	Title string `json:"title" form:"title" binding:"required"`
}