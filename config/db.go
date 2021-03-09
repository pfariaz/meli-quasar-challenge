package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/pfariaz/meli-quasar-challenge/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("sqlite3", "meli-quasar.db")

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&models.SatelliteMessage{})

	DB = database

}
