package repository

import (
	model "github.com/lancer2672/DandelionServer_Noti/db/model"
	"gorm.io/gorm"
)

type NotificationRepositoryInterface interface {
	AddNotification(notification model.Notification) error
	GetNotificationList() ([]model.Notification, error)
}

type NotificationRepo struct {
	DB *gorm.DB
}

func NewNotificationRepo(db *gorm.DB) NotificationRepositoryInterface {
	return &NotificationRepo{
		DB: db,
	}
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
