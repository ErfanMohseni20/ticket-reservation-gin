package customer

// Consolidated request types for Customer

type AddnewTicketRequest struct {
	Message string `json:"message" form:"message" binding:"required"`
	Title   string `json:"title" form:"title" binding:"required"`
}

type AddNewReplyToTicket struct {
	Message string `json:"message" form:"message" binding:"required"`
}

type ReserveSeat struct {
	BusID uint `json:"bus_id" form:"bus_id" binding:"required"`
}

type ReserveChangeStatus struct {
	ReserveID uint `json:"reserve_id" form:"reserve_id" binding:"required"`
	Status    string `json:"status" form:"status" binding:"required"`
}

type UpdateProfileRequest struct {
	FullName string `form:"full_name" binding:"omitempty,min=2,max=100"`
	Password string `form:"password" binding:"opiempty,min=6,max=100"`
}
