package model

import "github.com/uptrace/bun"

type Tables struct {
	bun.BaseModel `bun:"table:tables"`

	ID               int    `bun:",type:serial,autoincrement,pk"`
	TableNumber      string `bun:"table_number"`
	Capacity         string `bun:"capacity"`
	Status           string `bun:"status"`             // e.g., "available", "occupied", "reserved" , "unvailable"
	QrCodeIdentifier string `bun:"qr_code_identifier"` // Unique identifier for QR code
}
