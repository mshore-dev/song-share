package main

import (
	"log"

	"song-share/config"
	"song-share/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/handlebars/v2"
)

func main() {
	log.Println("song-share")

	// TODO: don't hardcode
	config.LoadConfig("./config.toml")

	app := fiber.New(fiber.Config{
		Views: handlebars.New("./assets/templates", ".hbs"),
	})

	app.Static("/assets", "./assets/static")

	// register all the other routes
	routes.RegisterRoutes(app)

	// it's go time.
	app.Listen(config.ActiveConfig.ListenAddr)
}
