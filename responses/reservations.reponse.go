package response

type ReservationResponse struct {
	ID              int    `json:"id"`
	TableID         int    `json:"table_id"`
	CustomerName    string `json:"customer_name"`
	CustomerPhone   string `json:"customer_phone"`
	ReservationTime string `json:"reservation_time"`
	NumberOfGuests  int    `json:"number_of_guests"`
	Status          string `json:"status"`
	Notes           string `json:"notes"`
	CreatedAt       int64  `json:"created_at"`
}
