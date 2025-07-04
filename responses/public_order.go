package response

type PublicOrderResponse struct {
	ID          int                  `json:"id"`
	TableID     int                  `json:"table_id"`
	TableNumber int                  `json:"table_number"`
	Status      string               `json:"status"`
	TotalAmount float64              `json:"total_amount"`
	Items       []OrderItemResponses `json:"items"`
	CreatedAt   int64                `json:"created_at"`
	UpdatedAt   int64                `json:"updated_at"`
	Message     string               `json:"message"`
}

type PublicMenuResponse struct {
	TableInfo TableInfo           `json:"table_info"`
	MenuItems []MenuItemResponses `json:"menu_items"`
}

type TableInfo struct {
	ID               int    `json:"id"`
	TableNumber      int    `json:"table_number"`
	QrCodeIdentifier string `json:"qr_code_identifier"`
	Status           string `json:"status"`
}
