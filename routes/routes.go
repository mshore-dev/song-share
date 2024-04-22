package routes

import (
	"song-share/routes/home"
	"song-share/routes/invite"
	"song-share/routes/list"
	"song-share/routes/user"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {

	home.RegisterRoutes(app)
	user.RegisterRoutes(app)
	list.RegisterRoutes(app)
	invite.RegisterRoutes(app)

}
