package admin


type TerminalResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	City CityResponse `json:"city"`
}
