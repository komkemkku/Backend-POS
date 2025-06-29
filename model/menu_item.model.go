package model

import "github.com/uptrace/bun"

type MenuItems struct {
	bun.BaseModel `bun:"table:menu_items"`

	ID          int     `bun:",type:serial,autoincrement,pk"`
	CategoryID  int     `bun:"category_id"`  // Foreign key to Categories table
	Name        string  `bun:"name"`         // Name of the menu item
	Description string  `bun:"description"`  // Description of the menu item
	Price       float64 `bun:"price"`        // Price of the menu item
	ImageURL    string  `bun:"image_url"`    // URL of the image for the menu
	IsAvailable bool    `bun:"is_available"` // Availability status of the menu item

	CreatedAt int64 `bun:"created_at"`
	UpdatedAt int64 `bun:"updated_at"`
}
