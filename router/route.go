package router

import (
	"fmt"

	"github.com/RifkyMuhamad/GolangAPIGateway/controllers"
	"github.com/gofiber/fiber/v2"
)

func NewRouter (app *fiber.App) {
	app.Get("/", controllers.HomeController)

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

	app.Get("/env", controllers.ENVController)
	app.Get("/string", controllers.StringController)

	app.Get("/belum-punya-pacar", func (c *fiber.Ctx) error {
		fmt.Println("Dia Belum punya cuy")
		return c.SendString("")
	})

	app.Get("/udah-punya-pacar", func (c *fiber.Ctx) error {
		fmt.Println("Cari yang lain aja")
		return c.SendString("")
	})
}