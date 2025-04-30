package database

import (
	"Chat-Application/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb(*gorm.DB, error) {
	dsn := "user:Happy8057@@tcp(127.0.0.1:3306)/Login?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	DB = db
	db.AutoMigrate(&models.User{})

	return db, nil

}
