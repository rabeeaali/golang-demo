package config

import (
	"fiber/app/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func InitialMigration() error {
	HOST := os.Getenv("DB_HOST")
	PORT := os.Getenv("DB_PORT")
	DBNAME := os.Getenv("DB_DATABASE")
	USER := os.Getenv("DB_USERNAME")
	DNS := fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Baghdad", HOST, USER, DBNAME, PORT)
	connection, err := gorm.Open(postgres.Open(DNS), &gorm.Config{})
	if err != nil {
		panic("could not connect to the database")
	}
	DB = connection
	AutoMigrateTables()
	return nil
}

func AutoMigrateTables() {
	DB.AutoMigrate(
		&models.User{},
		&models.ResetPassword{},
		&models.Post{},
	)
}
