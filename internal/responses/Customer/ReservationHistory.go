package customer

type ReservationHistory struct {
	ID                      uint   `json:"id"`
	OriginTerminalName      string `json:"origin_terminal_name"`
	DestinationTerminalName string `json:"destination_terminal_name"`
	Status                  string `json:"status"`
	ReservedAt              string `json:"reserved_at"`
	PurchasedAt 			string 	`json:"purchased_at"`
}
