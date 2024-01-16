package main

import (
	"log"
	"os"

	"github.com/RifkyMuhamad/GolangAPIGateway/router"
	"github.com/gofiber/fiber/v2"
)

func main() {

	// var db = database.OpenConnection()
	
	app := fiber.New()

	router.NewRouter(app)

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	log.Fatal(app.Listen("0.0.0.0:" + port))
}