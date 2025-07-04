package requests

type PublicOrderCreateRequest struct {
	QrCodeIdentifier string                   `json:"qr_code_identifier" binding:"required"`
	CustomerName     string                   `json:"customer_name"`
	CustomerPhone    string                   `json:"customer_phone"`
	Items            []PublicOrderItemRequest `json:"items" binding:"required,min=1"`
	Note             string                   `json:"note"`
}

type PublicOrderItemRequest struct {
	MenuItemID int `json:"menu_item_id" binding:"required"`
	Quantity   int `json:"quantity" binding:"required,min=1"`
}

