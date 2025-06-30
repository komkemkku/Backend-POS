package model

import "github.com/uptrace/bun"

type Tables struct {
	bun.BaseModel `bun:"table:tables"`

	ID               int    `bun:",pk,autoincrement"`
	TableNumber      int    `bun:"table_number"` // int เพื่อความเหมาะสม
	Capacity         int    `bun:"capacity"`
	Status           string `bun:"status"`             // e.g., "available", "occupied", "reserved", "unavailable"
	QrCodeIdentifier string `bun:"qr_code_identifier"` // Unique identifier for QR code
}
