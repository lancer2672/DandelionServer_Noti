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
	db, err := gorm.Open(postgres.Open(dbSource), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Connected to database")
	// Check if the type exists before creating it
	var count int64
	db.Raw("SELECT COUNT(*) FROM pg_type WHERE typname = 'notification_type'").Scan(&count)
	if count == 0 {
		db.Exec("CREATE TYPE notification_type AS ENUM ('chat', 'post', 'friend-request')")
	} else {
		log.Println("notification_type already exists")
	}
	if err := db.AutoMigrate(&model.Notification{}); err != nil {
		log.Fatalln("Auto migration failed:", err)
	}
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
