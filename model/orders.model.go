package model

import "github.com/uptrace/bun"

type Orders struct {
	bun.BaseModel `bun:"table:orders"`
	ID            int     `bun:",type:serial,autoincrement,pk"`
	TableID       int     `bun:"table_id"`
	StaffID       int     `bun:"staff_id"`
	Status        string  `bun:"status"`
	TotalAmount   float64 `bun:"total_amount"`

	CreateUnixTimestamp
	UpdateUnixTimestamp
}
