package model

import "github.com/uptrace/bun"

type Orders struct {
	bun.BaseModel `bun:"table:orders"`
	ID            int     `bun:",type:serial,autoincrement,pk"` // Unique
	TableID       int     `bun:"table_id"`                      // Foreign key to Table table
	StaffID       int     `bun:"staff_id"`                      // Foreign key to Staff table
	Status        string  `bun:"status"`                        // e.g., "pending", "completed", "cancelled"
	TotalAmout    float64 `bun:"total_amount"`                  // Total price of the order

	CreateUnixTimestamp
	UpdateUnixTimestamp

}
