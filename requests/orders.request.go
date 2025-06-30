package requests

type OrderRequest struct {
	Page   int64  `form:"page"`
	Size   int64  `form:"size"`
	Search string `form:"search"`
}

type OrderIdRequest struct {
	ID int `uri:"id"`
}

type OrderCreateRequest struct {
	TableID     int    `json:"table_id"`
	StaffID     int    `json:"staff_id"`
	Status      string `json:"status"`
	TotalAmount float64    `json:"total_amount"`
}

type OrderUpdateRequest struct {
	ID          int    `json:"id"`
	TableID     int    `json:"table_id"`
	StaffID     int    `json:"staff_id"`
	Status      string `json:"status"`
	TotalAmount float64    `json:"total_amount"`
}

type OrderDeleteRequest struct {
	ID int `json:"id"`
}
