package admin
type UsersResponse struct {
	ID uint `json:"id"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	CreatedAt string `json:"created_at"`
}