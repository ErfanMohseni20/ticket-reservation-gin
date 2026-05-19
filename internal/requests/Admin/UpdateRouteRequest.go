package admin

type UpdateRouteRequest struct {
	OriginTerminalId      uint   `json:"origin_terminal_id" form:"origin_terminal_id" binding:"omitempty"`
	DestinationTerminalId uint   `json:"destination_terminal_id" form:"destination_terminal_id" binding:"omitempty"`
	Duration              string `json:"duration" form:"duration" binding:"omitempty"`
	Distance              int    `json:"distance" form:"distance" binding:"omitempty"`
}