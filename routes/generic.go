package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pfariaz/meli-quasar-challenge/controllers"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func GenericRoutes(router *gin.Engine) {
	generic := router.Group("")
	{
		generic.GET("/", controllers.GetHealthcheck)

		url := ginSwagger.URL("/api-docs/doc.json")
		generic.GET("/api-docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	}
}
