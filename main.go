package main

import (
	"github.com/pfariaz/meli-quasar-challenge/config"
	"github.com/pfariaz/meli-quasar-challenge/routes"
	"github.com/pfariaz/meli-quasar-challenge/services"
)

func main() {
	r := routes.InitializeRoutes()
	services.LoadSatellites()
	config.ConnectDatabase()
	r.Run()
}
