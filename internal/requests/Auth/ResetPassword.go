package Auth


type ResetPasswordRequest struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"passowrd" binding:"required,min=6"`
	PasswordConfirmation string `json:"password_confirmation" form:"password_confirmation" binding:"required,eqfield=Password"`
}	
