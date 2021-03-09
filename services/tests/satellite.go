package services

import (
	"reflect"
	"testing"

	"github.com/pfariaz/meli-quasar-challenge/lib"
	"github.com/pfariaz/meli-quasar-challenge/models"
	"github.com/pfariaz/meli-quasar-challenge/services"
)

func TestGetSatellitesSuccessfull(t *testing.T) {

	var expectedSatelliteLocation map[string]models.PointLocation
	expectedSatelliteLocation[lib.KenobiSatelliteName] = models.PointLocation{X: -500, Y: -200}
	expectedSatelliteLocation[lib.SkywalkerSatelliteName] = models.PointLocation{X: 100, Y: -100}
	expectedSatelliteLocation[lib.SatoSatelliteName] = models.PointLocation{X: 500, Y: 100}

	services.LoadSatellites()

	if !reflect.DeepEqual(services.SatellitesLocation, expectedSatelliteLocation) {
		t.Error("The satellites from json are not equal to the expected")
	}
}

func TestGetSatellitesNames(t *testing.T) {

	expectedSatelliteNames := []string{"kenobi", "skywalker", "sato"}

	services.LoadSatellites()

	if !reflect.DeepEqual(services.GetSatellitesNames(), expectedSatelliteNames) {
		t.Error("The satellites from json are not equal to the expected")
	}
}
