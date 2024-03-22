package repository

import (
	"sync"

	"github.com/lancer2672/DandelionServer_Noti/db"
	model "github.com/lancer2672/DandelionServer_Noti/db/model"
	"gorm.io/gorm"
)

type NotificationRepositoryInterface interface {
	AddNotification(notification model.Notification) error
	GetNotificationList() ([]model.Notification, error)
	DeleteNotification(id int) error
}

var (
	instance *NotificationRepo
	once     sync.Once
)

func GetRepo() NotificationRepositoryInterface {
	once.Do(func() {
		instance = &NotificationRepo{
			DB: db.GetDB(),
		}
	})
	return instance
}

type NotificationRepo struct {
	DB *gorm.DB
}

func (r *NotificationRepo) AddNotification(notification model.Notification) error {
	result := r.DB.Create(&notification)
	return result.Error
}

func (r *NotificationRepo) GetNotificationList() ([]model.Notification, error) {
	var notifications []model.Notification
	result := r.DB.Find(&notifications)
	if result.Error != nil {
		return nil, result.Error
	}
	return notifications, nil
}
func (r *NotificationRepo) DeleteNotification(id int) error {
	result := r.DB.Delete(&model.Notification{}, id)
	return result.Error
}
