package model

import "github.com/uptrace/bun"

type Orders struct {
	bun.BaseModel `bun:"table:orders"`
	ID            int     `bun:",type:serial,autoincrement,pk"` // Unique
	TableID       int     `bun:"table_id"`                      // Foreign key to Table table
	StaffID       int     `bun:"staff_id"`                      // Foreign key to Staff table
	Status        string  `bun:"status"`                        // e.g., "pending", "completed", "cancelled"
	TotalAmout    float64 `bun:"total_amount"`                  // Total price of the order

	CreatedAt    int64 `bun:"created_at"`   // Timestamp for when the order was created
	Completed_at int64 `bun:"completed_at"` // Timestamp for when the order was completed

}
