package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pfariaz/meli-quasar-challenge/controllers"
)

func GenericRoutes(router *gin.Engine) {
	generic := router.Group("generic")
	{
		generic.GET("/", controllers.GetHealthcheck)

	}
}
