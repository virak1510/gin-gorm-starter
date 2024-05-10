package main

import (
	"myFirstGoGin/internal/routes"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=2906 dbname=gorm port=5432 TimeZone=Asia/Phnom_Penh"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// NOTE Optional config
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to get db instance")
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(10)

	router := routes.SetupRouter(db)
	router.Run(":8080")
}
