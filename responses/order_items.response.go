package response

type OrderItemResponses struct {
	ID           int     `json:"id"`
	OrderID      int     `json:"order_id"`
	MenuItemID   int     `json:"menu_item_id"`
	MenuName     string  `json:"menu_name"`
	Quantity     int     `json:"quantity"`
	PricePerItem float64 `json:"price_per_item"`
	SubTotal     float64 `json:"sub_total"`
	Notes        string  `json:"notes"`
	CreatedAt    int64   `json:"created_at"`
	UpdatedAt    int64   `json:"updated_at"`
}
