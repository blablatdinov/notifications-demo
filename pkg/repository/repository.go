package repository

import (
	"github.com/blablatdinov/notifications-demo"
	"github.com/jmoiron/sqlx"
)

type Notifications interface {
	Create(username string, NotificationString string) (int, error)
	GetNotifications(userId int) ([]notifications.Notification, error)
	DeleteNotification(username string, notificationId int) error
	GetNotificationsWithUsers() ([]notifications.NotificationWithUser, error)
}

type Authorization interface {
	CreateUser(notifications.User) (int, error)
	GetUser(username, password string) (notifications.User, error)
	GetUsers() ([]notifications.User, error)
}

type Repository struct {
	Notifications
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Notifications: NewNotificationsPostgres(db),
		Authorization: NewAuthPostgres(db),
	}
}
