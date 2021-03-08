package main

import (
	"github.com/pfariaz/meli-quasar-challenge/config"
	"github.com/pfariaz/meli-quasar-challenge/routes"
)

func main() {
	r := routes.InitializeRoutes()
	config.ConnectDatabase()
	r.Run()
}
