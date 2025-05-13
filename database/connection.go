package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBConnection() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	DBHost := os.Getenv("DB_HOST")
	PostgresUser := os.Getenv("POSTGRES_USER")
	PostgresPassword := os.Getenv("POSTGRES_PASSWORD")
	PostgresDB := os.Getenv("POSTGRES_DB")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai", DBHost, PostgresUser, PostgresPassword, PostgresDB)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error with connection: %v", err)
	}

	return db
}
