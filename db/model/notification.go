package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type JSON json.RawMessage

// Scan scan value into Jsonb, implements sql.Scanner interface
func (j *JSON) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	result := json.RawMessage{}
	err := json.Unmarshal(bytes, &result)
	*j = JSON(result)
	return err
}

// Value return json value, implement driver.Valuer interface
func (j JSON) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	return json.RawMessage(j).MarshalJSON()
}

type NotificationType string

const (
	chat          NotificationType = "chat"
	post          NotificationType = "post"
	friendRequest NotificationType = "friend-request"
)

func (st *NotificationType) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		*st = NotificationType(b)
	}
	return nil
}

func (st NotificationType) Value() (driver.Value, error) {
	return string(st), nil
}

type Notification struct {
	gorm.Model
	Type        NotificationType `gorm:"type:notification_type;default:'post'"`
	Description string           `gorm:"not null"`
	Title       string           `gorm:"default:''"`
	IsSeen      bool             `gorm:"default:false"`
	ReceiverID  uint             `gorm:"not null"`
	SenderID    uint             `gorm:"not null"`
	Payload     JSON             `gorm:"type:json"`
}
