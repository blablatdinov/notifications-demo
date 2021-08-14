package repository

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewPostgresDB() (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", "host=localhost port=5432 user=almazilaletdinov dbname=notifications sslmode=disable")
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, err
}
