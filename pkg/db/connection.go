package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() error {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbName, password)
	var err error
	db, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		return err
	}

	log.Println("Database connected...")
	return nil
}

func GetDB() *gorm.DB {
	return db
}
