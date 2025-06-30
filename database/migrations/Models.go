package migrations

import "Backend-POS/model"

func Models() []any {
	return []any{

		// (*model.Staff)(nil),
		// (*model.Categories)(nil),
		// (*model.Expenses)(nil),
		// (*model.MenuItems)(nil),
		// (*model.Orders)(nil),
		// (*model.OrderItems)(nil),
		// (*model.Payments)(nil),
		(*model.Reservations)(nil),
		// (*model.Tables)(nil),
	}
}

func RawBeforeQueryMigrate() []string {
	return []string{
		`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`,
	}
}

func RawAfterQueryMigrate() []string {
	return []string{}
}
