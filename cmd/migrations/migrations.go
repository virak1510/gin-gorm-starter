package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/virak1510/gin-gorm-starter/internal/schemas"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("migrating ....")
	err := godotenv.Load()
	if err != nil {
		panic("failed to load .env file")
	}
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&schemas.User{})

	fmt.Println("migrated successfully")
}
