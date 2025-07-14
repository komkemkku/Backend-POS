package model

import "github.com/uptrace/bun"

type Categories struct {
	bun.BaseModel `bun:"table:categories"`

	ID           int    `bun:",type:serial,autoincrement,pk"`
	Name         string `bun:"name"`
	Description  string `bun:"description"`
	DisplayOrder string `bun:"display_order"`
}
