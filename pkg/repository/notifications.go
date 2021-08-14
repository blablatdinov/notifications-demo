package repository

import (
	"fmt"

	"github.com/blablatdinov/notifications-demo"
	"github.com/jmoiron/sqlx"
)

type NotificationsPostgres struct {
	db *sqlx.DB
}

func NewNotificationsPostgres(db *sqlx.DB) *NotificationsPostgres {
	return &NotificationsPostgres{db: db}
}

func (r *NotificationsPostgres) GetNotifications(userId int) ([]notifications.Notification, error) {
	var result []notifications.Notification

	query := fmt.Sprintf("SELECT id, text FROM notifications WHERE user_id=%d", userId)
	if err := r.db.Select(&result, query); err != nil {
		return result, err
	}
	return result, nil
}

func (r *NotificationsPostgres) Create(username string, notificationText string) (int, error) {
	var id int

	userId, err := getUserId(*r.db, username)
	if err != nil {
		return 0, err
	}

	query := fmt.Sprintf("INSERT INTO notifications (text, user_id) values ('%s', %d) RETURNING id", notificationText, userId)
	row := r.db.QueryRow(query)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *NotificationsPostgres) DeleteNotification(username string, notificationId int) error {
	query := fmt.Sprintf("DELETE FROM notifications where id=%d", notificationId)
	_, err := r.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func getUserId(db sqlx.DB, username string) (int, error) {
	var id int
	query := fmt.Sprintf("select id from users where username='%s'", username)
	if err := db.Get(&id, query); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *NotificationsPostgres) GetNotificationsWithUsers() ([]notifications.NotificationWithUser, error) {
	var notifications []notifications.NotificationWithUser
	query := "select n.id, n.text, u.username from notifications as n inner join users as u on n.user_id = u.id"
	err := r.db.Select(&notifications, query)
	if err != nil {
		return nil, err
	}
	return notifications, nil
}
