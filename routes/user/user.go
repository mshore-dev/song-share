package user

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(app *fiber.App) {

	app.Get("/user/:user_id", routeGetUser)

	app.Post("/user/:user_id", routePostUpdateUser)

}

func routeGetUser(c *fiber.Ctx) error {

	return nil
}

func routePostUpdateUser(c *fiber.Ctx) error {

	return nil
}
