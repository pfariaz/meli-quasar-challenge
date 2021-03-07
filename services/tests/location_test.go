package services

import (
	"testing"

	"github.com/pfariaz/meli-quasar-challenge/services"
)

func TestGetLocationSuccessfull(t *testing.T) {

	pointX, pointY := services.GetLocation(100, 200, 300)

	if pointX == 0 && pointY == 0 {
		t.Error("Cannot determine the location")
	}
}

func TestGetLocationWithNoDistances(t *testing.T) {

	pointX, pointY := services.GetLocation()

	if pointX != 0 && pointY != 0 {
		t.Error("The method must return zero for both axes")
	}
}
