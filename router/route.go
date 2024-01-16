package router

import (
	"os"

	"github.com/RifkyMuhamad/GolangAPIGateway/controllers"
	"github.com/gofiber/fiber/v2"
)

func NewRouter (app *fiber.App) {
	app.Get("/", controllers.Welcome)

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
		mysqlLocal := os.Getenv("MYSQL_LOCAL")
		data := map[string]string{
			"port": test,
			"mysql": mysqlLocal,
		}
		return c.JSON(data)
	})

	app.Get("/string", func(c *fiber.Ctx) error {
		data := "Ini String"
		return c.SendString(data)
	})
}