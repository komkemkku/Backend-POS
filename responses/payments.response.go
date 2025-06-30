package response

type PaymentResponse struct {
	ID              int     `json:"id"`
	OrderID         int     `json:"order_id"`
	PaymentMethod   string  `json:"payment_method"`
	AmountPaid      float64 `json:"amount_paid"`
	TransactionTime string  `json:"transaction_time"`
}
