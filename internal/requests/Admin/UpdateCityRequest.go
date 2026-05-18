package admin

type UpdateCityRequest struct {
	Name string `json:"name" form:"name" binding:"required"`
}