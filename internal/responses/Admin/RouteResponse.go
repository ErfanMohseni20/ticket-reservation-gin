package admin

type RouteResponse struct {
	ID uint `json:"id"`
	OriginTerminal TerminalResponse `json:"origin_terminal"`
	DestinationTerminal TerminalResponse `json:"destination_terminal"`
	Duration string `json:"duration"`
	Distance int `json:"distance"`
}