package service

import (
	"github.com/blablatdinov/notifications-demo"
	"github.com/blablatdinov/notifications-demo/pkg/repository"
	"github.com/gorilla/websocket"
)

type Notifications interface {
	Create(username string, notificationText string) (int, error)
	GetNotifications(userId int) ([]notifications.Notification, error)
	DeleteNotification(username string, notificationId int) error
	GetChan() chan string
	ConnectUser(conn websocket.Conn, username string)
	GetNotificationsWithUsers() ([]notifications.NotificationWithUser, error)
	PingUsers()
}

type Authorization interface {
	CreateUser(user notifications.User) (int, error)
	GenerateToken(user notifications.User) (string, error)
	ParseToken(accessToken string) (int, string, error)
	GetUsers() ([]notifications.User, error)
}

type Service struct {
	Notifications
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Notifications: NewNotificationsService(repos.Notifications),
		Authorization: NewAuthService(repos.Authorization),
	}
}
