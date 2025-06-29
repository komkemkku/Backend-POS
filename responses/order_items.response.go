package response

type OrderItemResponses struct {
	ID           int    `json:"id"`
	OrderID      int    `json:"order_id"`
	MenuItemID   int    `json:"menu_item_id"`
	Quantity     int    `json:"quantity"`
	PricePerItem int    `json:"price_per_item"`
	Notes        string `json:"notes"`
}
