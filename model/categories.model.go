package model

import "github.com/uptrace/bun"

type Categories struct {
	bun.BaseModel `bun:"table:categories"`

	ID           int    `bun:",type:serial,autoincrement,pk"`
	Name         string `bun:"name"`          // Name of the category
	Description  string `bun:"description"`   // Description of the category
	DisplayOrder string `bun:"display_order"` // Order in which the category should be displayed
}
