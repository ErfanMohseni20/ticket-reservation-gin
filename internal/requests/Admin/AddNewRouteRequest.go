package admin

type AddNewRouteRequest struct {
	OriginTerminalId      uint   `json:"origin_terminal_id" form:"origin_terminal_id" binding:"required"`
	DestinationTerminalId uint   `json:"destination_terminal_id" form:"destination_terminal_id" binding:"required"`
	Duration              string `json:"duration" form:"duration" binding:"required"`
	Distance              int    `json:"distance" form:"distance" binding:"required"`
}