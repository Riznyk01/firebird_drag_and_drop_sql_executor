package repository

import (
	"database/sql"
	"time"
)

type Viewing interface {
	UpdateDBCorrectionDate(currentTime time.Time) error
	GetDBCorrectionDate() (time.Time, error)
}

type Repository struct {
	Viewing
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Viewing: NewFirebirdClient(db),
	}
}
