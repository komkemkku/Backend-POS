package model

import "github.com/uptrace/bun"

type Expenses struct {
	bun.BaseModel `bun:"table:expenses"`
	ID            int     `bun:",type:serial,autoincrement,pk"` // Unique
	Decscription  string  `bun:"description"`                   // Description of the expense
	Amount        float64 `bun:"amount"`                        // Amount of the expense
	Category      string  `bun:"category"`                      // Category of the expense (e.g
	ExpenseDate   int64   `bun:"expense_date"`                  // Date of the expense in Unix timestamp format
	StaffID       int     `bun:"staff_id"`                      // Foreign key to Staff table

	CreateUnixTimestamp
}
