package customer

type ProfileResponse struct {
	Fullname  string `json:"fullname"`
	Username  string `json:"username"`
	CreatedAt string `json:"created_at"`
	AvatarURL string `json:"avatar_url"`
}
