package controllers

import "github.com/gofiber/fiber/v2"

func Welcome(c *fiber.Ctx) error {
	data := map[string]string{
		"documentation": "https://github.com/RifkyMuhamad/GolangAPIGateway",
		"message": "Dyone Strankers use Go Fiber",
	}
	return c.JSON(data)
}