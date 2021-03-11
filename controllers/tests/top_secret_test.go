package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/pfariaz/meli-quasar-challenge/config"
	"github.com/pfariaz/meli-quasar-challenge/lib/constants"
	"github.com/pfariaz/meli-quasar-challenge/lib/helpers"
	"github.com/pfariaz/meli-quasar-challenge/models/schemas"
	"github.com/pfariaz/meli-quasar-challenge/routes"
)

func TestPostProcessMessageLocation(t *testing.T) {
	helpers.LoadSatellitesTesting()

	t.Run("Returns decoded message and position successfully", func(t *testing.T) {
		expectedResponse := schemas.ShipResponseSchema{
			Position: schemas.PositionResponseSchema{
				X: -487.2859125,
				Y: 1557.014225,
			},
			Message: "este es un mensaje secreto",
		}

		requestPayload := schemas.ProcessFullMessageRequestSchema{
			Satellites: []schemas.SatelliteRequestSchema{
				{
					Name:     constants.KenobiSatelliteName,
					Distance: 100.0,
					Message:  []string{"este", "", "", "mensaje", ""},
				},
				{
					Name:     constants.SkywalkerSatelliteName,
					Distance: 115.5,
					Message:  []string{"", "es", "", "", "secreto"},
				},
				{
					Name:     constants.SatoSatelliteName,
					Distance: 142.7,
					Message:  []string{"este", "", "un", "", ""},
				},
			},
		}
		payloadBuf := new(bytes.Buffer)
		json.NewEncoder(payloadBuf).Encode(requestPayload)

		w := httptest.NewRecorder()
		router := routes.InitializeRoutes()

		req, _ := http.NewRequest("POST", "/api/v1/topsecret/", payloadBuf)
		req.Header.Add("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		byteValue, _ := ioutil.ReadAll(w.Body)
		var response schemas.ShipResponseSchema
		json.Unmarshal(byteValue, &response)

		if w.Code != http.StatusOK {
			t.Fail()
		}

		if !reflect.DeepEqual(response, expectedResponse) {
			t.Fail()
		}
	})

	t.Run("Respond with error because sends less than 3 required satellites", func(t *testing.T) {
		expectedResponse := schemas.HTTPError{Error: "We need to provide the information of the 3 known satellites (kenobi, skywalker and sato)"}

		requestPayload := schemas.ProcessFullMessageRequestSchema{
			Satellites: []schemas.SatelliteRequestSchema{
				{
					Name:     constants.KenobiSatelliteName,
					Distance: 100.0,
					Message:  []string{"este", "", "", "mensaje", ""},
				},
				{
					Name:     constants.SkywalkerSatelliteName,
					Distance: 115.5,
					Message:  []string{"", "es", "", "", "secreto"},
				},
			},
		}
		payloadBuf := new(bytes.Buffer)
		json.NewEncoder(payloadBuf).Encode(requestPayload)

		w := httptest.NewRecorder()
		router := routes.InitializeRoutes()

		req, _ := http.NewRequest("POST", "/api/v1/topsecret/", payloadBuf)
		req.Header.Add("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		byteValue, _ := ioutil.ReadAll(w.Body)
		var response schemas.HTTPError
		json.Unmarshal(byteValue, &response)

		if w.Code == http.StatusOK {
			t.Fail()
		}

		if !reflect.DeepEqual(response, expectedResponse) {
			t.Fail()
		}
	})

	t.Run("Respond with error because sends unknown satellites", func(t *testing.T) {
		expectedResponse := schemas.HTTPError{Error: "we cannot determine the location since we receive unknown satellites (known satellites are kenobi, skywalker and sato)"}

		requestPayload := schemas.ProcessFullMessageRequestSchema{
			Satellites: []schemas.SatelliteRequestSchema{
				{
					Name:     "meli",
					Distance: 100.0,
					Message:  []string{"este", "", "", "mensaje", ""},
				},
				{
					Name:     "darthvader",
					Distance: 115.5,
					Message:  []string{"", "es", "", "", "secreto"},
				},
				{
					Name:     "mandalorian",
					Distance: 142.7,
					Message:  []string{"este", "", "un", "", ""},
				},
			},
		}
		payloadBuf := new(bytes.Buffer)
		json.NewEncoder(payloadBuf).Encode(requestPayload)

		w := httptest.NewRecorder()
		router := routes.InitializeRoutes()

		req, _ := http.NewRequest("POST", "/api/v1/topsecret/", payloadBuf)
		req.Header.Add("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		byteValue, _ := ioutil.ReadAll(w.Body)
		var response schemas.HTTPError
		json.Unmarshal(byteValue, &response)

		if w.Code == http.StatusOK {
			t.Fail()
		}

		if !reflect.DeepEqual(response, expectedResponse) {
			t.Fail()
		}
	})
}

func TestProcessPartialMessageLocation(t *testing.T) {
	helpers.LoadSatellitesTesting()
	config.ConnectDatabase()

	t.Run("Save successfully satellite info", func(t *testing.T) {
		requestPayload := schemas.ProcessSplitMessageRequestSchema{
			Distance: 100.0,
			Message:  []string{"este", "", "", "mensaje", ""},
		}
		payloadBuf := new(bytes.Buffer)
		json.NewEncoder(payloadBuf).Encode(requestPayload)

		w := httptest.NewRecorder()
		router := routes.InitializeRoutes()

		req, _ := http.NewRequest("POST", "/api/v1/topsecret_split/kenobi", payloadBuf)
		req.Header.Add("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		if w.Code != http.StatusNoContent {
			t.Fail()
		}
	})

	t.Run("Respond with error because sends unknown satellite", func(t *testing.T) {
		expectedResponse := schemas.HTTPError{Error: "We need you to provide us with the information of the 3 known satellites (kenobi, skywalker or sato)"}

		requestPayload := schemas.ProcessSplitMessageRequestSchema{
			Distance: 100.0,
			Message:  []string{"este", "", "", "mensaje", ""},
		}
		payloadBuf := new(bytes.Buffer)
		json.NewEncoder(payloadBuf).Encode(requestPayload)

		w := httptest.NewRecorder()
		router := routes.InitializeRoutes()

		req, _ := http.NewRequest("POST", "/api/v1/topsecret_split/mercadolibre", payloadBuf)
		req.Header.Add("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		byteValue, _ := ioutil.ReadAll(w.Body)
		var response schemas.HTTPError
		json.Unmarshal(byteValue, &response)

		if w.Code == http.StatusOK {
			t.Fail()
		}

		if !reflect.DeepEqual(response, expectedResponse) {
			t.Fail()
		}
	})

	helpers.DestroyTestingDatabaseFile()
}

func TestGetPartialMessageLocation(t *testing.T) {
	helpers.LoadSatellitesTesting()
	config.ConnectDatabase()

	t.Run("Get location and decoded message successfully because it has 3 satellites info ", func(t *testing.T) {

		cleaner := helpers.CleanDB()
		defer cleaner()

		helpers.LoadSatellitesIntoDB()

		expectedResponse := schemas.ShipResponseSchema{
			Position: schemas.PositionResponseSchema{
				X: -487.2859125,
				Y: 1557.014225,
			},
			Message: "este es un mensaje secreto",
		}

		w := httptest.NewRecorder()
		router := routes.InitializeRoutes()

		req, _ := http.NewRequest("GET", "/api/v1/topsecret_split/", nil)
		router.ServeHTTP(w, req)

		byteValue, _ := ioutil.ReadAll(w.Body)
		var response schemas.ShipResponseSchema
		json.Unmarshal(byteValue, &response)

		fmt.Println(w.Code)

		if w.Code != http.StatusOK {
			t.Fail()
		}

		if !reflect.DeepEqual(response, expectedResponse) {
			t.Fail()
		}
	})

	t.Run("Return error because we dont have info of 3 satellites", func(t *testing.T) {

		cleaner := helpers.CleanDB()
		defer cleaner()

		expectedResponse := schemas.HTTPError{Error: "To determine exact message and position, we need the information of more satellites"}

		w := httptest.NewRecorder()
		router := routes.InitializeRoutes()

		req, _ := http.NewRequest("GET", "/api/v1/topsecret_split/", nil)
		router.ServeHTTP(w, req)

		byteValue, _ := ioutil.ReadAll(w.Body)
		var response schemas.HTTPError
		json.Unmarshal(byteValue, &response)

		if w.Code != http.StatusBadRequest {
			t.Fail()
		}

		if !reflect.DeepEqual(response, expectedResponse) {
			t.Fail()
		}
	})

	helpers.DestroyTestingDatabaseFile()
}
