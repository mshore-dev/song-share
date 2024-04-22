package list

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(app *fiber.App) {

	app.Get("/list", routeGetList)
	app.Get("/list/:list_id", routeGetListOne)
	app.Get("/list/:list_id/add", routeGetListAddSong)

	app.Post("/list", routePostListCreate)
	app.Post("/list/:list_id/add", routePostListAddSong)
	// app.Post("/list/:list_id/invite", routePostListInvite) // ?
	app.Post("/list/:list_id/:song_id/comment", routePostListCommentSong)
	app.Post("/list/:list_id/:song_id/vote", routePostListVoteSong)

	app.Delete("/list/:list_id", routeDeleteList)
	app.Delete("/list/:list_id/:song_id", routeDeleteListSong)
	app.Delete("/list/:list_id/:song_id/:comment_id", routeDeleteListCommentSong)

}

func routeGetList(c *fiber.Ctx) error {

	c.Render("list", fiber.Map{
		"title": "Demo List",
	})

	return nil
}

func routeGetListOne(c *fiber.Ctx) error {

	return nil
}

func routePostListCreate(c *fiber.Ctx) error {

	return nil
}

func routeDeleteList(c *fiber.Ctx) error {

	return nil
}

func routeGetListAddSong(c *fiber.Ctx) error {

	return nil
}

func routePostListAddSong(c *fiber.Ctx) error {

	return nil
}

// func routePostListInvite(c *fiber.Ctx) error {

// 	return nil
// }

func routeDeleteListSong(c *fiber.Ctx) error {

	return nil
}

func routePostListVoteSong(c *fiber.Ctx) error {

	return nil
}

func routePostAddSongComment(c *fiber.Ctx) error {

	return nil
}

func routePostListCommentSong(c *fiber.Ctx) error {

	return nil
}

func routeDeleteListCommentSong(c *fiber.Ctx) error {

	return nil
}
