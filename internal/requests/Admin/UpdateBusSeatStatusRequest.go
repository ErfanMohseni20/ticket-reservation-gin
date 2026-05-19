package admin

type UpdateBusSeatStatusRequest struct {
	SeatID uint `json:"seat_id" form:"seat_id" binding:"required"`
	BusID uint `json:"bus_id" form:"bus_id" binding:"required"`
	Status string `json:"status" form:"status" binding:"required"`
}