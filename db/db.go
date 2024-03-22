package db

import (
	"log"
	"sync"

	"github.com/lancer2672/DandelionServer_Noti/db/model"
	"github.com/lancer2672/DandelionServer_Noti/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	instance *gorm.DB
	once     sync.Once
)

func initDB(dbSource string) *gorm.DB {
	// Khởi tạo kết nối cơ sở dữ liệu ở đây
	db, err := gorm.Open(postgres.Open(dbSource), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Connected to database")
	db.Exec("CREATE TYPE notification_type AS ENUM ('chat', 'post', 'friend-request')")
	db.AutoMigrate(&model.Notification{})
	return db
}
func GetDB() *gorm.DB {
	once.Do(func() {
		config, err := utils.LoadConfig(".")
		if err != nil {
			log.Fatal("Cannot load config", err)
		}
		//if db not connected, init using default db_source
		instance = initDB(config.DB_SOURCE)
	})
	return instance
}
