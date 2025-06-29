package requests

type ReservationRequest struct {
	Page   int64  `form:"page"`
	Size   int64  `form:"size"`
	Search string `form:"search"`
}

type ReservationIdRequest struct {
	ID int `uri:"id"`
}

type ReservationCreateRequest struct {
	TableID         int    `json:"table_id"`
	CustomerName    string `json:"customer_name"`
	CustomerPhone   string `json:"customer_phone"`
	ReservationTime string `json:"reservation_time"`
	NumberOfGuests  int    `json:"number_of_guests"`
	Ststus          string `json:"status"`
	Notes           string `json:"notes"`
}

type ReservationUpdateRequest struct {
	ID              int    `json:"id"`
	TableID         int    `json:"table_id"`
	CustomerName    string `json:"customer_name"`
	CustomerPhone   string `json:"customer_phone"`
	ReservationTime string `json:"reservation_time"`
	NumberOfGuests  int    `json:"number_of_guests"`
	Status          string `json:"status"`
	Notes           string `json:"notes"`
}

type ReservationDeleteRequest struct {
	ID int `json:"id"`
}
