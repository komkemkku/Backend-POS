package response

type OrderResponse struct {
	ID          int                  `json:"id"`
	TableID     int                  `json:"table_id"`
	StaffID     int                  `json:"staff_id"`
	Status      string               `json:"status"`
	TotalAmount float64              `json:"total_amount"`
	CreatedAt   int64                `json:"created_at"`
	CompletedAt int64                `json:"completed_at"`
	OrderItems  []OrderItemResponses `json:"order_items,omitempty"`
}
