package main

import (
	"log"
	"os"

	"github.com/RifkyMuhamad/GolangAPIGateway/router"
	"github.com/gofiber/fiber/v2"
)

func main() {

	// database.OpenConnection()

	app := fiber.New()

	router.NewRouter(app)

	app.Get("/env", func(c *fiber.Ctx) error {
		test := os.Getenv("TEST_ENV")
		data := map[string]string{
			"port": test,
		}
		return c.JSON(data)
	})

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	log.Fatal(app.Listen("0.0.0.0:" + port))
}