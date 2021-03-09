package controllers

import (
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pfariaz/meli-quasar-challenge/config"
	"github.com/pfariaz/meli-quasar-challenge/lib"
	"github.com/pfariaz/meli-quasar-challenge/models"
	"github.com/pfariaz/meli-quasar-challenge/models/schemas"
	"github.com/pfariaz/meli-quasar-challenge/services"
)

// ProcessMessageLocation godoc
// @Summary Process Message given info by all satellites
// @Param satellites body schemas.ProcessFullMessageRequestSchema true "Add satellites info"
// @Accept  json
// @Produce  json
// @Success 200 {object} schemas.ShipResponseSchema
// @Failure 400 {object} schemas.HTTPError
// @Router /topsecret/ [post]
func ProcessMessageLocation(c *gin.Context) {
	var request schemas.ProcessFullMessageRequestSchema

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	qtySatellites := len(request.Sattelites)
	if qtySatellites < 3 || qtySatellites > 3 {
		c.JSON(http.StatusBadRequest, schemas.HTTPError{Error: "We need to provide the information of the 3 known sattelites (kenobi, skywalker and sato)"})
		return
	}

	var givenSattelites []string
	var distances []float64
	var messages [][]string
	for _, sattelite := range request.Sattelites {
		givenSattelites = append(givenSattelites, sattelite.Name)
		distances = append(distances, sattelite.Distance)
		messages = append(messages, sattelite.Message)
	}

	if !reflect.DeepEqual(givenSattelites, services.GetSatellitesNames()) {
		c.JSON(http.StatusBadRequest, schemas.HTTPError{Error: "we cannot determine the location since we receive unknown satellites (known satellites are kenobi, skywalker and sato)"})
		return
	}

	message := services.GetMessage(messages...)
	xPosition, yPosition := services.GetLocation(distances...)

	var messageLocationShip = schemas.ShipResponseSchema{
		Position: schemas.PositionResponseSchema{
			X: xPosition,
			Y: yPosition,
		},
		Message: message,
	}

	c.JSON(http.StatusOK, messageLocationShip)
}

// ProcessPartialMessageLocation godoc
// @Summary Process Message given info by each satellite
// @Param satellite_name path string true "string enums" Enums(kenobi, skywalker, sato)
// @Param satellite body schemas.ProcessSplitMessageRequestSchema true "Add satellite info"
// @Accept  json
// @Produce  json
// @Success 204
// @Failure 400 {object} schemas.HTTPError
// @Router /topsecret_split/{satellite_name} [post]
func ProcessPartialMessageLocation(c *gin.Context) {
	var request schemas.ProcessSplitMessageRequestSchema
	satelliteName := c.Param("satellite_name")
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, schemas.HTTPError{Error: err.Error()})
		return
	}

	if !lib.Contains(services.GetSatellitesNames(), satelliteName) {
		c.JSON(http.StatusBadRequest, schemas.HTTPError{Error: "We need you to provide us with the information of the 3 known satellites (kenobi, skywalker or sato)"})
		return
	}

	// Create satelliteInfo
	messageJoined := strings.Join(request.Message, ",")
	satellite := models.SatelliteMessage{Name: satelliteName, Distance: request.Distance, Message: messageJoined}
	config.DB.Where("name = ?", satelliteName).First(&satellite)

	if satellite.ID != 0 {
		satellite.Distance = request.Distance
		satellite.Message = messageJoined
		config.DB.Save(&satellite)
	} else {
		config.DB.Create(&satellite)
	}

	c.Writer.WriteHeader(http.StatusNoContent)
}

// GetPartialMessageLocation godoc
// @Summary Get location given info from 3 satellites
// @Param satellite_name path string true "string enums" Enums(kenobi, skywalker, sato)
// @Accept  json
// @Produce  json
// @Success 200 {object} schemas.ShipResponseSchema
// @Failure 400 {object} schemas.HTTPError
// @Router /topsecret_split/{satellite_name} [get]
func GetPartialMessageLocation(c *gin.Context) {
	var satellitesCreated []models.SatelliteMessage
	config.DB.Find(&satellitesCreated)

	if len(satellitesCreated) < len(services.GetSatellitesNames()) {
		c.JSON(http.StatusBadRequest, schemas.HTTPError{Error: "To determine exact message and position, we need the information of more satellites"})
		return
	}
	var distances []float64
	var messages [][]string

	for _, sattelite := range satellitesCreated {
		distances = append(distances, sattelite.Distance)
		messages = append(messages, strings.Split(sattelite.Message, ","))
	}

	message := services.GetMessage(messages...)
	xPosition, yPosition := services.GetLocation(distances...)

	var messageLocationShip = schemas.ShipResponseSchema{
		Position: schemas.PositionResponseSchema{
			X: xPosition,
			Y: yPosition,
		},
		Message: message,
	}

	c.JSON(http.StatusOK, messageLocationShip)
}
