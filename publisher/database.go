package publisher

import (
	"fmt"
	"log"
	"os"
	"sqs-blog/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartDatabase() *gorm.DB {
	DB_HOST := os.Getenv("DB_HOST")
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")
	DB_PORT := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	if err := db.AutoMigrate(&models.Post{}); err != nil {
		log.Fatalln(err)
	}

	// log.Println("Database connected")

	return db
}
