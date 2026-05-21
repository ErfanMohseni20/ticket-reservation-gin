package admin

type TicketResponse struct {
	ID uint `json:"id"`
	Title string `json:"title"`
	CreatorName string `json:"creator_name"`
	Status string `json:"status"`
	CreatedAt string `json:"created_at"`
}