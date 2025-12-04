package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func main() {
	app := fiber.New()
	app.Use(healthcheck.New())
	app.Use(logger.New(logger.Config{Format: "[${ip}]:${port} ${status} - ${method} ${path}\n"}))
	app.Get("/metrics", monitor.New())

	log.Fatal(app.Listen(":3000"))
}
