package admin

// Consolidated request types for Admin

type AddNewCityRequest struct {
	Name string `json:"name" form:"name" binding:"required"`
}

type UpdateCityRequest struct {
	Name string `json:"name" form:"name" binding:"required"`
}

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

type AddNewRouteRequest struct {
	OriginTerminalId      uint   `json:"origin_terminal_id" form:"origin_terminal_id" binding:"required"`
	DestinationTerminalId uint   `json:"destination_terminal_id" form:"destination_terminal_id" binding:"required"`
	Duration              string `json:"duration" form:"duration" binding:"required"`
	Distance              int    `json:"distance" form:"distance" binding:"required"`
}

type UpdateRouteRequest struct {
	OriginTerminalId      uint   `json:"origin_terminal_id" form:"origin_terminal_id" binding:"omitempty"`
	DestinationTerminalId uint   `json:"destination_terminal_id" form:"destination_terminal_id" binding:"omitempty"`
	Duration              string `json:"duration" form:"duration" binding:"omitempty"`
	Distance              int    `json:"distance" form:"distance" binding:"omitempty"`
}

type AddNewTerminalRequest struct {
	Name string `json:"name" form:"name" binding:"required"`
	CityId uint `json:"city_id" form:"city_id" binding:"required"`
}

type UpdateTerminalRequest struct {
	Name string `json:"name" form:"name" binding:"omitempty"`
	CityId uint `json:"city_id" form:"city_id" binding:"omitempty,numric"`
}

type UpdateBusSeatStatusRequest struct {
	SeatID uint `json:"seat_id" form:"seat_id" binding:"required"`
	BusID uint `json:"bus_id" form:"bus_id" binding:"required"`
	Status string `json:"status" form:"status" binding:"required"`
}

type AddNewUserReqeust struct {
	FullName string `json:"fullname" form:"fullname" binding:"required"`
	UserName string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required,min=6"`
	PasswordConfirmation string `json:"password_confirmation" form:"password_confirmation" binding:"requried,eqfield=password"`
}

type UpdateUserRequest struct {
	FullName             string `json:"fullname" form:"fullname" binding:"omitempty"`
	UserName             string `json:"username" form:"username" binding:"omitempty"`
	Password             string `json:"password" form:"password" binding:"omitempty,min=6"`
	PasswordConfirmation string `json:"password_confirmation" form:"password_confirmation" binding:"requried,eqfield=password"`
}

type AddNewReplyToTicket struct {
	Message string `json:"message" form:"message" binding:"required"`
}
