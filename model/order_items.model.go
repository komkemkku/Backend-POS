package model

import "github.com/uptrace/bun"

type OrderItems struct {
	bun.BaseModel `bun:"table:order_items"`
	ID            int     `bun:",type:serial,autoincrement,pk"`
	OrderID       int     `bun:"order_id"`
	MenuItemID    int     `bun:"menu_item_id"`
	Quantity      int     `bun:"quantity"`
	PricePerItem  float64 `bun:"price_per_item"`
	Notes         string  `bun:"notes"`

	CreateUnixTimestamp
	UpdateUnixTimestamp
}
