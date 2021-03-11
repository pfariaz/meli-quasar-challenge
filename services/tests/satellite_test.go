package services

import (
	"reflect"
	"testing"

	"github.com/pfariaz/meli-quasar-challenge/lib/constants"
	"github.com/pfariaz/meli-quasar-challenge/lib/helpers"
	"github.com/pfariaz/meli-quasar-challenge/models"
	"github.com/pfariaz/meli-quasar-challenge/services"
)

func TestGetSatellitesSuccessfull(t *testing.T) {

	var expectedSatelliteLocation = map[string]models.PointLocation{
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

	helpers.LoadSatellitesTesting()

	if !reflect.DeepEqual(services.SatellitesLocation, expectedSatelliteLocation) {
		t.Error("The satellites from json are not equal to the expected")
	}
}

func TestGetSatellitesNames(t *testing.T) {

	expectedSatelliteNames := []string{"kenobi", "skywalker", "sato"}

	helpers.LoadSatellitesTesting()

	if !helpers.Equal(services.GetSatellitesNames(), expectedSatelliteNames) {
		t.Error("The satellites from json are not equal to the expected")
	}
}
