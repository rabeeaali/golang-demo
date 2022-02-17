package main

import (
	"fiber/config"
	"fiber/routes"
)

func main() {
	// Load our environment variables
	config.LoadEnv()
	// init database config
	config.InitialMigration()
	// init route
	routes.SetupRoutes()
}
