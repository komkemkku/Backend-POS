package model

import "github.com/uptrace/bun"

type Payments struct {
	bun.BaseModel   `bun:"table:payments"`
	ID              int     `bun:",type:serial,autoincrement,pk"` // Unique
	OrderID         int     `bun:"order_id"`                      // Foreign key to Orders table
	PaymentMethod   string  `bun:"payment_method"`                // e.g., "cash", "credit_card", "mobile_payment"
	AmountPaid      float64 `bun:"amount_paid"`                   // Amount paid
	TransactionTime int64   `bun:"transaction_time"`              // Timestamp for when the payment was made

}
