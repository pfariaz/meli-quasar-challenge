package main

import (
	"github.com/pfariaz/meli-quasar-challenge/config"
	"github.com/pfariaz/meli-quasar-challenge/docs"
	"github.com/pfariaz/meli-quasar-challenge/lib/helpers"
	"github.com/pfariaz/meli-quasar-challenge/routes"
	"github.com/pfariaz/meli-quasar-challenge/services"
)

func main() {

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Quasar Fire Meli Test API"
	docs.SwaggerInfo.Description = "This is MercadoLibre test API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = helpers.GetBaseURL()
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"https"}

	r := routes.InitializeRoutes()
	services.LoadSatellites()
	config.ConnectDatabase()
	r.Run()
}
