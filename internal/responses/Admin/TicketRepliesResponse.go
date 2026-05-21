package admin

type TicketRepliesResponse struct { 
	ID uint `json:"id"`
	Message string `json:"message"`
	SenderRole string `json:"sender_role"`
	CreatedAt string `json:"created_at"`
}