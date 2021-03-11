package config

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/pfariaz/meli-quasar-challenge/models"
)

var DB *gorm.DB

func GetDBName() string {
	environment := os.Getenv("GIN_MODE")
	db_url := "meli-quasar.db"
	if environment == gin.TestMode {
		db_url = "meli-testing-quasar.db"
	}

	return db_url
}

func ConnectDatabase() {
	database, err := gorm.Open("sqlite3", GetDBName())

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&models.SatelliteMessage{})

	DB = database
}
