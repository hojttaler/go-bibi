package database

import (
	"fmt"

	"discord-bot/src/config"
	"discord-bot/src/database/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase(connectLink string) {
	db, err := gorm.Open(mysql.Open(config.Config.Database.Link), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())

		return
	}

	db.AutoMigrate(&models.SalaryHistory{})

	DB = db
}
