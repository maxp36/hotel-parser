package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/maxp36/hotel-parser/app"
)

type repository struct {
	DB *sqlx.DB
}

// NewRepository creates an object that represent the v2.PostgresRepository interface
func NewRepository(db *sqlx.DB) app.Repository {
	return &repository{
		DB: db,
	}
}
