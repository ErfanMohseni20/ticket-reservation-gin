package admin

type BusResponse struct {
	ID               uint   `json:"id"`
	RouteID          uint   `json:"route_id"`
	DepartureTime    string `json:"departure_time"`
	ArrivalTime      string `json:"arrival_time"`
	Capacity         int    `json:"capacity"`
	Price            int    `json:"price"`
	BusType          string `json:"bus_type"`
	Corporation      string `json:"corporation"`
	SuperCorporation string `json:"super_corporation"`
	ServiceNumber    string `json:"service_number"`
	IsVIP            bool   `json:"is_vip"`
	
}
