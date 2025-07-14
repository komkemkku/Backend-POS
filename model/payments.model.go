package model

import "github.com/uptrace/bun"

type Payments struct {
	bun.BaseModel   `bun:"table:payments"`
	ID              int     `bun:",type:serial,autoincrement,pk"`
	OrderID         int     `bun:"order_id"`
	PaymentMethod   string  `bun:"payment_method"`
	AmountPaid      float64 `bun:"amount_paid"`
	TransactionTime int64   `bun:"transaction_time"`
}
