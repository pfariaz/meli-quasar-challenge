package services

import (
	"testing"

	"github.com/pfariaz/meli-quasar-challenge/services"
)

func TestGetMessageSuccessfull(t *testing.T) {

	var fragmentsMessageOne = []string{"I", "want", "", "", "", ""}
	var fragmentsMessageTwo = []string{"", "", "to", "work", "", ""}
	var fragmentsMessageThree = []string{"", "", "", "", "in", ""}
	var fragmentsMessageFour = []string{"", "", "", "", "", "MercadoLibre"}

	expectedMessage := "I want to work in MercadoLibre"
	messageReceived := services.GetMessage(fragmentsMessageOne, fragmentsMessageTwo, fragmentsMessageThree, fragmentsMessageFour)

	if messageReceived != expectedMessage {
		t.Errorf("The received messsage must be equal to %s", expectedMessage)
	}
}

func TestGetMessageEmptyBecauseNoFragmentsGiven(t *testing.T) {

	messageReceived := services.GetMessage()

	if len(messageReceived) > 0 {
		t.Error("The received messsage must be empty")
	}
}
