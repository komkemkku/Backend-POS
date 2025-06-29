package requests

type PaymentRequest struct {
	Page   int64  `form:"page"`
	Size   int64  `form:"size"`
	Search string `form:"search"`
}

type PaymentIdRequest struct {
	ID int `uri:"id"`
}

type PaymentCreateRequest struct {
	OrderID         int    `json:"order_id"`
	PaymentMethod   string `json:"payment_method"`
	AmountPaid      int    `json:"amount_paid"`
	TransactionTime string `json:"transaction_time"`
}

type PaymentUpdateRequest struct {
	ID              int    `json:"id"`
	OrderID         int    `json:"order_id"`
	PaymentMethod   string `json:"payment_method"`
	AmountPaid      int    `json:"amount_paid"`
	TransactionTime string `json:"transaction_time"`
}

type PaymentDeleteRequest struct {
	ID int `json:"id"`
}
