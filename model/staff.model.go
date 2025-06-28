package model

import "github.com/uptrace/bun"

type Staff struct {
	bun.BaseModel `bun:"table:staff"`

	ID           int    `bun:",type:serial,autoincrement,pk"`
	UserName     string `bun:"username"`
	PasswordHash string `bun:"password_hash"`
	FullName     string `bun:"full_name"`
	Role         string `bun:"role"`

	CreatedAt int64 `bun:"created_at"`
	UpdatedAt int64 `bun:"updated_at"`
}
