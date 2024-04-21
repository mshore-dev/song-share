package main

import (
	"log"

	"song-share/config"

	"github.com/gofiber/fiber/v2"
)

func main() {
	log.Println("song-share")

	// TODO: don't hardcode
	config.LoadConfig("./config.toml")

	app := fiber.New()

	app.Listen(config.ActiveConfig.ListenAddr)
}
