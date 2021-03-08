package controllers

import (
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pfariaz/meli-quasar-challenge/config"
	"github.com/pfariaz/meli-quasar-challenge/models"
	"github.com/pfariaz/meli-quasar-challenge/models/schemas"
	"github.com/pfariaz/meli-quasar-challenge/services"
)

func GetSatellitesNames() []string {
	return []string{"kenobi", "skywalker", "sato"}
}

func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func ProcessMessageLocation(c *gin.Context) {
	var request schemas.ProcessFullMessageRequestSchema

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	qtySatellites := len(request.Sattelites)
	if qtySatellites < 3 || qtySatellites > 3 {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "We need to provide the information of the 3 known sattelites (kenobi, skywalker and sato)"})
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

	if !reflect.DeepEqual(givenSattelites, GetSatellitesNames()) {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "we cannot determine the location since we receive unknown satellites (known satellites are kenobi, skywalker and sato)"})
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

func ProcessPartialMessageLocation(c *gin.Context) {
	var request schemas.ProcessSplitMessageRequestSchema
	satelliteName := c.Param("satellite_name")
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !Contains(GetSatellitesNames(), satelliteName) {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "We need you to provide us with the information of the 3 known satellites (kenobi, skywalker or sato)"})
		return
	}

	// Create satellite
	messageJoined := strings.Join(request.Message, ",")
	satellite := models.Satellite{Name: satelliteName, Distance: request.Distance, Message: messageJoined}
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

func GetPartialMessageLocation(c *gin.Context) {
	var satellitesCreated []models.Satellite
	config.DB.Find(&satellitesCreated)

	if len(satellitesCreated) < len(GetSatellitesNames()) {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "To determine exact message and position, we need the information of more satellites"})
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
