package home

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	app.Get("/", routeGetHome)
}

func routeGetHome(c *fiber.Ctx) error {

	log.Printf("hello from routeGetHome")

	c.Render("home", fiber.Map{
		"title": "Home",
	})

	return nil
}
