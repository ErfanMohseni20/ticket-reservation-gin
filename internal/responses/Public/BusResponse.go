package public

import (
	bus "github.com/ErfanMohseni20/ticket-reservation-gin/internal/responses/Admin"
)

type BusResponse struct {
	ID                      uint            `json:"id"`
	OriginTerminalName      string          `json:"origin_terminal_name"`
	DestinationTerminalName string          `json:"destination_terminal_name"`
	AvailableSeatsCount     int64           `json:"available_seats_count"`
	BusData                 bus.BusResponse `json:"bus_data"`
}
