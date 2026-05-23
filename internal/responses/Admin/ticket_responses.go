package admin

type TicketResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	CreatorName string `json:"creator_name"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
}

type TicketRepliesResponse struct {
	ID         uint   `json:"id"`
	Message    string `json:"message"`
	SenderRole string `json:"sender_role"`
	CreatedAt  string `json:"created_at"`
}
