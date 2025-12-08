package config

import (
	"fmt"
	"log"
      "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "root:@tcp(127.0.0.1:3306)/gin-crud-api?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	DB = db
	fmt.Println("✅ Database connected!")
}

func AutoMigrate(models ...interface{}) {
	DB.AutoMigrate(models...)
	fmt.Println("✅ Database migrated!")
}
