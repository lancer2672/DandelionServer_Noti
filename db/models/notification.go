package models

type Notification struct {
	Id    int    `json:"id" gorm:"primaryKey"`
	Title string `json:"title"`
}
