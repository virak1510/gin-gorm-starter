package main

import (
	"fmt"
	"myFirstGoGin/internal/schemas"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("migrating ....")
	dsn := "host=localhost user=postgres password=2906 dbname=gorm port=5432 TimeZone=Asia/Phnom_Penh"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&schemas.User{})

	fmt.Println("migrated successfully")
}
