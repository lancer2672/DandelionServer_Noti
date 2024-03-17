package services

import (
	"github.com/lancer2672/DandelionServer_Noti/db/models"
	"github.com/lancer2672/DandelionServer_Noti/utils"
	"gorm.io/gorm"
)

type NotificationService struct {
	db *gorm.DB
}

func New(db *gorm.DB) *NotificationService {
	return &NotificationService{
		db,
	}
}

func (s *NotificationService) CreateNotifcation(notification models.Notification) {
	result := s.db.Create(&notification)
	if result.Error != nil {
		utils.FailOnError(result.Error, "Create notification failed")
	}
}
