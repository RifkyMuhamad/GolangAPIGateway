package controllers

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func ENVController(c *fiber.Ctx) error {
	test := os.Getenv("TEST_ENV")
		mysqlLocal := os.Getenv("MYSQL_LOCAL")
		data := map[string]string{
			"port": test,
			"mysql": mysqlLocal,
		}
		return c.JSON(data)
}