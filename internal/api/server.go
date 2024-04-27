package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nesar-ahmed/goecommerce/config"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()

	app.Listen(config.ServerPort)
}