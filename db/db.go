package db

import (
	"log"

	"github.com/lancer2672/DandelionServer_Noti/db/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(dbSource string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dbSource), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Connected to database")
	db.Exec("CREATE TYPE notification_type AS ENUM ('chat', 'post', 'friend-request')")
	db.AutoMigrate(&model.Notification{})

	return db
}
