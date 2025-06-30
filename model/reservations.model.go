package model

import "github.com/uptrace/bun"

type Reservations struct {
	bun.BaseModel   `bun:"table:reservations"`
	ID              int    `bun:",pk,autoincrement"`
	TableID         int    `bun:"table_id"`
	CustomerName    string `bun:"customer_name"`
	CustomerPhone   string `bun:"customer_phone"`
	ReservationTime int64  `bun:"reservation_time"`
	NumberOfGuests  int    `bun:"number_of_guests"`
	Status          string `bun:"status"`
	Notes           string `bun:"notes"` // Additional notes for the reservation

	CreateUnixTimestamp
}
