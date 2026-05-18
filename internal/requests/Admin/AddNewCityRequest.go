package admin

type AddNewCityRequest struct {
	Name string `json:"name" form:"name" binding:"required"`
}