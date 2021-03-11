package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pfariaz/meli-quasar-challenge/routes"
)

func TestGetHealthcheckSuccessfull(t *testing.T) {

	type HealthcheckBody struct {
		Time string `json:"time"`
	}

	w := httptest.NewRecorder()

	router := routes.InitializeRoutes()

	req, _ := http.NewRequest("GET", "/", nil)

	router.ServeHTTP(w, req)

	byteValue, _ := ioutil.ReadAll(w.Body)
	var healthcheckBody HealthcheckBody
	json.Unmarshal(byteValue, &healthcheckBody)

	if len(healthcheckBody.Time) == 0 {
		t.Fail()
	}

	if w.Code == http.StatusNotFound {
		t.Fail()
	}

}
