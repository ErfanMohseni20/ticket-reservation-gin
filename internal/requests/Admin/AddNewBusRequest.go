package admin

type AddNewBusRequest struct {
	RouteID          int    `json:"route_id" form:"route_id" binding:"required"`
	DepartureTime    string `json:"departure_time" form:"departure_time" binding:"required"`
	ArrivalTime      string `json:"arrival_time" form:"arrival_time" binding:"required"`
	Capacity         int    `json:"capacity" form:"capacity" binding:"required"`
	Price            int    `json:"price" form:"price" binding:"required"`
	BusType          string `json:"bus_type" form:"bus_type" binding:"required"`
	Corporation      string `json:"corporation" form:"corporation" binding:"required"`
	SuperCorporation string `json:"super_corporation" form:"super_corporation" binding:"required"`
	ServiceNumber    string `json:"service_number" form:"service_number" binding:"required"`
	IsVIP            bool   `json:"is_vip" form:"is_vip" binding:"required"`
}
