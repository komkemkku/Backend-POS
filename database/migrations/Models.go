package migrations

import "Backend-POS/models"


func Models() []any {
	return []any{

		(*models.Users)(nil),
		(*models.Admins)(nil),

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
