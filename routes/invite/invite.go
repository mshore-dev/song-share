package invite

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(app *fiber.App) {

	app.Get("/invite/:invite_code", routeGetInvite)

	app.Delete("/invite/:invite_code", routeDeleteInvite)

}

func routeGetInvite(c *fiber.Ctx) error {

	return nil
}

func routeDeleteInvite(c *fiber.Ctx) error {

	return nil
}
