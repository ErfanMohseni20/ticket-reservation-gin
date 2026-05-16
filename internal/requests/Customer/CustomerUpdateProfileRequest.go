package customer

import "mime/multipart"


type UpdateProfileRequest struct {
	FullName string                `form:"full_name" binding:"omitempty,min=2,max=100"`
	Avatar   *multipart.FileHeader `form:"avatar"`
}