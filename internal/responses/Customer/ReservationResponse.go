package customer

type ReservationResponse struct {
	ID         uint   `json:"id"`
	BusID      uint   `json:"bus_id"`
	SeatNumber int    `json:"seat_number"`
	Status     string `json:"status"`
	ReservedAt string `json:"reserved_at"`
}