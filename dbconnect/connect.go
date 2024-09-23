package dbconnect

import (
	"fmt"
	"goapi/config"
	"goapi/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	config := config.GetConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DB_USER,
		config.DB_PASS,
		config.DB_HOST,
		config.DB_PORT,
		config.DB_NAME)

    var err error
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %s", err.Error()))
	}
	fmt.Println("database connected!")

	fmt.Println("migrating schema")

	DB.AutoMigrate(models.User{})
    DB.AutoMigrate(models.Payment{})
    DB.AutoMigrate(models.Category{})

	fmt.Println("schema migrated")
}
