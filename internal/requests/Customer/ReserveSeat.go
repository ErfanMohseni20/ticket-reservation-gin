package customer

type ReserveSeat struct {
	BusID uint `json:"bus_id" form:"bus_id" binding:"required"`
}