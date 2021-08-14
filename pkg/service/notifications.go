package service

import (
	"github.com/blablatdinov/notifications-demo"
	"github.com/blablatdinov/notifications-demo/pkg/repository"
	"github.com/gorilla/websocket"
)

type NotificationsService struct {
	repo repository.Notifications
	Hub  notifications.Hub
}

func NewNotificationsService(repo repository.Notifications) *NotificationsService {
	return &NotificationsService{
		repo: repo,
		Hub:  notifications.NewHub(make(chan string, 5)),
	}
}

func (s *NotificationsService) ConnectUser(conn websocket.Conn, username string) {
	s.Hub.ConnectUser(conn, username)
}

func (s *NotificationsService) GetChan() chan string {
	return s.Hub.MessageChan
}

func (s *NotificationsService) Create(username string, notificationText string) (int, error) {
	err := s.Hub.Send(username, notificationText)
	if err == nil {
		return 0, nil
	}
	id, err := s.repo.Create(username, notificationText)
	if err != nil {
		return id, err
	}
	return id, nil
}

func (s *NotificationsService) GetNotifications(userId int) ([]notifications.Notification, error) {
	result, err := s.repo.GetNotifications(userId)
	if err != nil {
		return make([]notifications.Notification, 0), err
	}
	return result, nil
}

func (s *NotificationsService) DeleteNotification(username string, notificationId int) error {
	err := s.repo.DeleteNotification(username, notificationId)
	if err != nil {
		return err
	}
	return nil
}

func (s *NotificationsService) GetNotificationsWithUsers() ([]notifications.NotificationWithUser, error) {
	notifications, err := s.repo.GetNotificationsWithUsers()
	if err != nil {
		return nil, err
	}
	return notifications, nil
}

func (s *NotificationsService) PingUsers() {
	s.Hub.PingUsers()
}
