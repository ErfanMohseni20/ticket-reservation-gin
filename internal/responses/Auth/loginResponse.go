package Auth

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Message      string `json:"message"`
	User         struct {
		Username string `json:"username"`
		FullName string `json:"full_name"`
	} `json:"user"`
}
