package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(dbSource string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dbSource), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Connected to database")
	// db.AutoMigrate(&models.Book{})

	return db
}
