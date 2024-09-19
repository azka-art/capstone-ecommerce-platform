package main

import (
	"ecommerce-platform/backend/config"
	"ecommerce-platform/backend/routes"
	"log"
)

func main() {
	// Initialize the database
	config.ConnectDB()

	// Set up the router
	r := routes.SetupRouter()

	// Start the server
	log.Fatal(r.Run(":8080"))
}
