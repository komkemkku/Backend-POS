package requests

type OrderItemRequest struct {
	Page   int64  `form:"page"`
	Size   int64  `form:"size"`
	Search string `form:"search"`
}

type OrderItemIdRequest struct {
	ID int `uri:"id"`
}

type OrderItemCreateRequest struct {
	OrderID      int    `json:"order_id"`
	MenuItemID   int    `json:"menu_item_id"`
	Quantity     int    `json:"quantity"`
	PricePerItem int    `json:"price_per_item"`
	Notes        string `json:"notes"`
}

type OrderItemUpdateRequest struct {
	ID           int    `json:"id"`
	OrderID      int    `json:"order_id"`
	MenuItemID   int    `json:"menu_item_id"`
	Quantity     int    `json:"quantity"`
	PricePerItem int    `json:"price_per_item"`
	Notes        string `json:"notes"`
}

type OrderItemDeleteRequest struct {
	ID int `json:"id"`
}
