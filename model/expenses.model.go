package model

import "github.com/uptrace/bun"

type Expenses struct {
	bun.BaseModel `bun:"table:expenses"`
	ID            int     `bun:",type:serial,autoincrement,pk"`
	Description   string  `bun:"description"`
	Amount        float64 `bun:"amount"`
	Category      string  `bun:"category"`
	ExpenseDate   string  `bun:"expense_date"`
	StaffID       int     `bun:"staff_id"`

	CreateUnixTimestamp
}
