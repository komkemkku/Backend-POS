package model

import "github.com/uptrace/bun"

type OrderItems struct {
	bun.BaseModel `bun:"table:order_items"`
	ID            int     `bun:",type:serial,autoincrement,pk"` // Unique
	OrderID       int     `bun:"order_id"`                      // Foreign key to Orders table
	MenuItemID    int     `bun:"menu_item_id"`                  // Foreign key to MenuItem table
	Quantity      int     `bun:"quantity"`                      // Quantity of the menu item ordered
	PricePerItem  float64 `bun:"price_per_item"`                // Price of the menu item at the reservation time
	Notes         string  `bun:"notes"`                         // Additional notes for the order item

}
