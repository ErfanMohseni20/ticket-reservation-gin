package admin

type AddNewTerminalRequest struct {
	Name string `json:"name" form:"name" binding:"required"`
	CityId uint `json:"city_id" form:"city_id" binding:"required"`
}