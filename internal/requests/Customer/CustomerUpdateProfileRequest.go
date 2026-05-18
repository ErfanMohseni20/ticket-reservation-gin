package customer
type UpdateProfileRequest struct {
	FullName string `form:"full_name" binding:"omitempty,min=2,max=100"`
	Password string `form:"password" binding:"opiempty,min=6,max=100"`
}
