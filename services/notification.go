package services

import (
	"log"
	"sync"

	"github.com/lancer2672/DandelionServer_Noti/db/model"
	"github.com/lancer2672/DandelionServer_Noti/db/repository"
)

type NotificationService struct {
	Repo repository.NotificationRepositoryInterface
}

var (
	instance *NotificationService
	once     sync.Once
)

func GetService() *NotificationService {
	once.Do(func() {
		instance = &NotificationService{
			Repo: repository.GetRepo(),
		}
	})
	return instance
}

// @Summary Add notification
// @Description add a new notification
// @Tags notifications
// @Accept  json
// @Produce  json
// @Param notification body model.Notification true "Add new notification"
// @Success 200 {object} map[string]interface{}
// @Router /notification [post]
func (s *NotificationService) AddNotification(notification model.Notification) error {
	err := s.Repo.AddNotification(notification)
	if err != nil {
		log.Fatal("Create notification failed", err)
		return err
	}
	return nil
}

// @Summary Get notifications
// @Description get all notifications
// @Tags notifications
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]interface{}
// @Router /notification [get]
func (s *NotificationService) GetNotificationList() ([]model.Notification, error) {
	list, err := s.Repo.GetNotificationList()
	if err != nil {
		log.Fatal("Get notification list failed", err)
		return nil, err
	}
	return list, nil
}
