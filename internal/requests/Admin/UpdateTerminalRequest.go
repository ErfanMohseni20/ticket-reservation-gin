package admin

type UpdateTerminalRequest struct {
	Name string `json:"name" form:"name" binding:"omitempty"`
	CityId uint `json:"city_id" form:"city_id" binding:"omitempty,numric"`
}