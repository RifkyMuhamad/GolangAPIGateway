package main

import (
	"os"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	// var db = database.OpenConnection()
	
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		data := map[string]string{
			"message": "Dyone Strankers use Go Fiber",
			"documentation": "https://github.com/RifkyMuhamad/GolangAPIGateway",
		}
		return c.JSON(data)
	})

	app.Get("/categories", func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{
			"message": "Ini Get Categories",
		})
	})

	app.Get("/categories/:categoriesId", func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{
			"message": "Ini Get Categories CategoriesId",
		})
	})

	app.Post("/categories", func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{
			"message": "Ini Post Categories",
		})
	})

	app.Put("/categories/:categoryId", func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{
			"message": "Ini Put Categories CategoriesId",
		})
	})

	app.Delete("/categories/:categoryId", func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{
			"message": "Ini Delete Categories CategoriesId",
		})
	})

	app.Get("/env", func(c *fiber.Ctx) error {
		test := os.Getenv("TEST_ENV")
		data := map[string]string{
			"port": test,
		}
		return c.JSON(data)
	})

	app.Get("/string", func(c *fiber.Ctx) error {
		data := "Ini String"
		return c.SendString(data)
	})

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	log.Fatal(app.Listen("0.0.0.0:" + port))
}