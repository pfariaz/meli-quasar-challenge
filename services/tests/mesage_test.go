package services

import (
	"testing"

	"github.com/pfariaz/meli-quasar-challenge/services"
)

func TestGetMessageSuccessfull(t *testing.T) {

	var fragments_message_one = []string{"I", "want", "", "", "", ""}
	var fragments_message_two = []string{"", "", "to", "work", "", ""}
	var fragments_message_three = []string{"", "", "", "", "in", ""}
	var fragments_message_four = []string{"", "", "", "", "", "MercadoLibre"}

	expected_message := "I want to work in MercadoLibre"
	message_received := services.GetMessage(fragments_message_one, fragments_message_two, fragments_message_three, fragments_message_four)

	if message_received != expected_message {
		t.Errorf("The received messsage must be equal to %s", expected_message)
	}
}

func TestGetMessageEmptyBecauseNoFragmentsGiven(t *testing.T) {

	message_received := services.GetMessage()

	if len(message_received) > 0 {
		t.Error("The received messsage must be empty")
	}
}
