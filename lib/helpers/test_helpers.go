package helpers

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/pfariaz/meli-quasar-challenge/config"
	"github.com/pfariaz/meli-quasar-challenge/lib/constants"
	"github.com/pfariaz/meli-quasar-challenge/models"
	"github.com/pfariaz/meli-quasar-challenge/services"
)

func LoadSatellitesTesting() {
	services.SatellitesLocation = map[string]models.PointLocation{
		constants.KenobiSatelliteName: {
			X: -500,
			Y: -200,
		},
		constants.SkywalkerSatelliteName: {
			X: 100,
			Y: -100,
		},
		constants.SatoSatelliteName: {
			X: 500,
			Y: 100,
		},
	}
}

func DestroyTestingDatabaseFile() {
	os.Remove(config.GetDBName())
}

func LoadSatellitesIntoDB() {
	config.DB.Create(&models.SatelliteMessage{Name: constants.KenobiSatelliteName, Distance: 100.0, Message: "este,,,mensaje,"})
	config.DB.Create(&models.SatelliteMessage{Name: constants.SkywalkerSatelliteName, Distance: 115.5, Message: ",es,,,secreto"})
	config.DB.Create(&models.SatelliteMessage{Name: constants.SatoSatelliteName, Distance: 142.7, Message: "este,,un,,"})
}

func CleanDB() func() {
	type entity struct {
		table   string
		keyname string
		key     interface{}
	}
	var entries []entity
	hookName := "cleanupHook"

	config.DB.Callback().Create().After("gorm:create").Register(hookName, func(scope *gorm.Scope) {
		entries = append(entries, entity{table: scope.TableName(), keyname: scope.PrimaryKey(), key: scope.PrimaryKeyValue()})
	})
	return func() {
		// Remove the hook once we're done
		defer config.DB.Callback().Create().Remove(hookName)
		// Find out if the current db object is already a transaction
		_, inTransaction := config.DB.CommonDB().(*sql.Tx)
		tx := config.DB
		if !inTransaction {
			tx = config.DB.Begin()
		}
		// Loop from the end. It is important that we delete the entries in the
		// reverse order of their insertion
		for i := len(entries) - 1; i >= 0; i-- {
			entry := entries[i]
			fmt.Printf("Deleting entities from '%s' table with key %v\n", entry.table, entry.key)
			tx.Table(entry.table).Where(entry.keyname+" = ?", entry.key).Delete("")
		}

		if !inTransaction {
			tx.Commit()
		}
	}
}
