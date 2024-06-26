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
		"title":       "Demo List",
		"description": "A really cool demo playlist!",
		"songs": []fiber.Map{
			{
				"link":          "https://www.youtube.com/watch?v=3nlSDxvt6JU",
				"title":         "Snail's House - Pixel Galaxy (Official MV)",
				"comment":       "A really cool song I liked!",
				"created_at":    "11:09 PM - 1 Jan 2016",
				"score":         7,
				"comment_count": 1,
			},
			{
				"link":          "https://www.youtube.com/watch?v=uZf8a_tE00Q",
				"title":         "Pixel Fantasy - Pixel Galaxy (Official MV)",
				"comment":       "Another catchy song!",
				"created_at":    "12:09 PM - 2 Jan 2016",
				"score":         8,
				"comment_count": 2,
			},
			{
				"link":          "https://www.youtube.com/watch?v=n4U20g40jZc",
				"title":         "Supercell - Snowflakes (Official MV)",
				"comment":       "A beautiful song!",
				"created_at":    "1:09 PM - 3 Jan 2016",
				"score":         -2,
				"comment_count": 3,
			},
			{
				"link":          "https://www.youtube.com/watch?v=a_uG_h0v34o",
				"title":         "Pixel Heart - Pixel Galaxy (Official MV)",
				"comment":       "A touching song!",
				"created_at":    "2:09 PM - 4 Jan 2016",
				"score":         10,
				"comment_count": 4,
			},
		},
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
