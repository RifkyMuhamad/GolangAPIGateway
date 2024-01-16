package controllers

import "github.com/gofiber/fiber/v2"

func StringController(c *fiber.Ctx) error {
	data := "Ini String"
	return c.SendString(data)
}
