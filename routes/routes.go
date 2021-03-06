package routes

import (
	"github.com/gin-gonic/gin"
)

func InitializeRoutes() *gin.Engine {
	router := gin.New()

	GenericRoutes(router)
	TopSecretRoutes(router)

	return router
}
