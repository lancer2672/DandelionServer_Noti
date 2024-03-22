package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
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
	Chat          NotificationType = "chat"
	Post          NotificationType = "post"
	FriendRequest NotificationType = "friend-request"
)

func (nt *NotificationType) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		*nt = NotificationType(v)
	case string:
		*nt = NotificationType(v)
	default:
		return errors.New("unsupported type for NotificationType")
	}
	fmt.Println("VALUE", *nt)
	return nil
}
func (nt NotificationType) Value() (driver.Value, error) {
	return string(nt), nil
}

type Notification struct {
	ID          uuid.UUID        `json:"id" gorm:"type:uuid;primary_key"`
	CreatedAt   time.Time        `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt   gorm.DeletedAt   `gorm:"index" json:"-"`
	Type        NotificationType `json:"type" gorm:"type:notification_type;default:'post'"`
	Description string           `json:"description" gorm:"not null"`
	Title       string           `json:"title" gorm:"default:''"`
	IsSeen      bool             `json:"isSeen" gorm:"default:false"`
	ReceiverID  uint             `json:"receiverId" gorm:"not null"`
	SenderID    uint             `json:"senderId" gorm:"not null"`
	Payload     json.RawMessage  `json:"payload"`
}

func (u *Notification) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	//other check
	return
}
