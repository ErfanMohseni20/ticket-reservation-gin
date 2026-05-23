package Auth

// Consolidated request types for Auth

type ResetPasswordRequest struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"passowrd" binding:"required,min=6"`
	PasswordConfirmation string `json:"password_confirmation" form:"password_confirmation" binding:"required,eqfield=Password"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Fullname string `json:"fullname" binding:"required"`
}
