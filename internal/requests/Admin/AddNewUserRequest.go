package admin

type AddNewUserReqeust struct {
	FullName string `json:"fullname" form:"fullname" binding:"required"`
	UserName string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required,min=6"`
	PasswordConfirmation string `json:"password_confirmation" form:"password_confirmation" binding:"requried,eqfield=password"`
}