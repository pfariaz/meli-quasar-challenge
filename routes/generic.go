package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pfariaz/meli-quasar-challenge/controllers"
	"github.com/pfariaz/meli-quasar-challenge/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func GenericRoutes(router *gin.Engine) {

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Quasar Fire Meli Test API"
	docs.SwaggerInfo.Description = "This is MercadoLibre test API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "quasarmeli.swagger.io"
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	generic := router.Group("")
	{
		generic.GET("/", controllers.GetHealthcheck)
		generic.GET("/api-docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	}
}
