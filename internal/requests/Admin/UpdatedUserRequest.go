package admin

type UpdateUserRequest struct {
	FullName             string `json:"fullname" form:"fullname" binding:"omitempty"`
	UserName             string `json:"username" form:"username" binding:"omitempty"`
	Password             string `json:"password" form:"password" binding:"omitempty,min=6"`
	PasswordConfirmation string `json:"password_confirmation" form:"password_confirmation" binding:"requried,eqfield=password"`
}
