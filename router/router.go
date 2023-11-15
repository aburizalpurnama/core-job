package router

import (
	"core-users-job/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func InitRoutes(cfg fiber.Config, h handler.ReportHandler) *fiber.App {
	app := fiber.New(cfg)
	app.Use(recover.New())
	app.Use(cors.New())
	app.Use(logger.New(logger.ConfigDefault))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON("Welcome")
	})

	report := app.Group("/reports")
	report.Post("open-account", h.GenerateOpenAccountReport)

	return app
}
