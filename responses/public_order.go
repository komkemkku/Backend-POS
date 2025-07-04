package response

type PublicOrderResponse struct {
	ID          int                  `json:"id"`
	TableID     int                  `json:"table_id"`
	Status      string               `json:"status"`
	TotalAmount float64              `json:"total_amount"`
	Items       []OrderItemResponses `json:"items"`
	CreatedAt   int64                `json:"created_at"`
	Message     string               `json:"message"`
}
