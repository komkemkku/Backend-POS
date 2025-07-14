package model

import "github.com/uptrace/bun"

type Tables struct {
	bun.BaseModel `bun:"table:tables"`

	ID               int    `bun:",pk,autoincrement"`
	TableNumber      int    `bun:"table_number"`
	Capacity         int    `bun:"capacity"`
	Status           string `bun:"status"`
	QrCodeIdentifier string `bun:"qr_code_identifier"`
}
