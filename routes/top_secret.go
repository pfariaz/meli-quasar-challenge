package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pfariaz/meli-quasar-challenge/controllers"
)

func TopSecretRoutes(router *gin.Engine) {
	v1 := router.Group("api/v1")
	{
		v1.POST("/topsecret/", controllers.ProcessMessageLocation)
		v1.POST("/topsecret_split/:satellite_name", controllers.ProcessPartialMessageLocation)
		v1.GET("/topsecret_split/", controllers.GetPartialMessageLocation)

	}
}
