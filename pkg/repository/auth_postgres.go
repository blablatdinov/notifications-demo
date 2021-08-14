package repository

import (
	"github.com/blablatdinov/notifications-demo"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) GetUsers() ([]notifications.User, error) {
	var users []notifications.User
	query := "select id, username from users"
	err := r.db.Select(&users, query)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *AuthPostgres) CreateUser(user notifications.User) (int, error) {
	var id int
	query := "INSERT INTO users (username, password_hash) values ($1, $2) RETURNING id"
	row := r.db.QueryRow(query, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (notifications.User, error) {
	var user notifications.User
	query := "select id from users where username=$1 and password_hash=$2"
	err := r.db.Get(&user, query, username, password)
	return user, err
}
