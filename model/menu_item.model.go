package model

import "github.com/uptrace/bun"

type MenuItems struct {
	bun.BaseModel `bun:"table:menu_items"`

	ID          int     `bun:",type:serial,autoincrement,pk"`
	CategoryID  int     `bun:"category_id"`
	Name        string  `bun:"name"`
	Description string  `bun:"description"`
	Price       float64 `bun:"price"`
	ImageURL    string  `bun:"image_url"`
	IsAvailable bool    `bun:"is_available"`

	CreateUnixTimestamp
	UpdateUnixTimestamp
}
