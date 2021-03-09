package services

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/pfariaz/meli-quasar-challenge/lib"
	"github.com/pfariaz/meli-quasar-challenge/models"
)

var SatellitesLocation map[string]models.PointLocation

func LoadSatellites() {

	type SatelliteJSON struct {
		Name string  `json:"name"`
		X    float64 `json:"x"`
		Y    float64 `json:"y"`
	}
	type SatellitesJSON struct {
		Satellites []SatelliteJSON `json:"satellites"`
	}

	jsonFile, err := os.Open(lib.SatellitesPath)
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var satellitesJSON SatellitesJSON

	json.Unmarshal(byteValue, &satellitesJSON)

	SatellitesLocation = make(map[string]models.PointLocation)
	for _, satellite := range satellitesJSON.Satellites {
		SatellitesLocation[satellite.Name] = models.PointLocation{X: satellite.X, Y: satellite.Y}
	}
}

func GetSatellitesNames() []string {

	var satellitesNames []string
	for key := range SatellitesLocation {
		satellitesNames = append(satellitesNames, key)
	}
	return satellitesNames
}
