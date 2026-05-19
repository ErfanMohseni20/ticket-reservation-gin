package admin

type UpdateBusRequest struct {
	RouteID          *int    `json:"route_id" form:"route_id" binding:"omitempty"`
	DepartureTime    string `json:"departure_time" form:"departure_time" binding:"omitempty"`
	ArrivalTime      string `json:"arrival_time" form:"arrival_time" binding:"omitempty"`
	Capacity         int    `json:"capacity" form:"capacity" binding:"omitempty"`
	Price            int    `json:"price" form:"price" binding:"omitempty"`
	BusType          string `json:"bus_type" form:"bus_type" binding:"omitempty"`
	Corporation      string `json:"corporation" form:"corporation" binding:"omitempty"`
	SuperCorporation string `json:"super_corporation" form:"super_corporation" binding:"omitempty"`
	ServiceNumber    string `json:"service_number" form:"service_number" binding:"omitempty"`
	IsVIP            bool   `json:"is_vip" form:"is_vip" binding:"omitempty"`
}
