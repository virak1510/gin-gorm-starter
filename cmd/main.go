package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/virak1510/gin-gorm-starter/internal/routes"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("failed to load .env file")
	}
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// NOTE Optional config
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to get db instance")
	}
	sqlDB.SetMaxIdleConns(4)
	sqlDB.SetMaxOpenConns(32)

	router := routes.SetupRouter(db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}
