package model

import "github.com/uptrace/bun"

type Reservations struct {
	bun.BaseModel   `bun:"table:reservations"`
	ID              int    `bun:",type:serial,autoincrement,pk"` // Unique
	CustomerName    string `bun:"customer_name"`                 // Name of the customer making the reservation
	CustomerPhone   string `bun:"customer_phone"`                // Phone number of the customer
	ReservationTime int64  `bun:"reservation_time"`              // Timestamp for the reservation
	NumberOfGuests  int    `bun:"number_of_guests"`              // Number of guests for the reservation
	Status          string `bun:"status"`                        // Status of the reservation (e.g., "pending", "confirmed", "cancelled")
	Notes           string `bun:"notes"`                         // Additional notes for the reservation

	CreateUnixTimestamp
}
