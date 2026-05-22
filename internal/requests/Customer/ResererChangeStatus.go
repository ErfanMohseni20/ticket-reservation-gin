package customer

type ReserveChangeStatus struct {
	ReserveID uint `json:"reserve_id" form:"reserve_id" binding:"required"`
	Status string `json:"status" form:"status" binding:"required"`
}