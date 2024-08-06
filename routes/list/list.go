package list

import (
	"context"
	"fmt"
	"song-share/database"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {

	app.Get("/list", routeGetList)
	app.Get("/list/:list_id", routeGetList)
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

	listID, err := c.ParamsInt("list_id")
	if err != nil || listID == 0 {
		c.Redirect("/list")
	}

	list, err := database.Queries.GetList(context.Background(), int64(listID))
	if err != nil {
		return c.SendString(fmt.Sprintf("failed to query database for list:\r\n\r\n %v", err))
	}

	songs, err := database.Queries.GetSongsByList(context.Background(), int64(listID))
	if err != nil {
		return c.SendString(fmt.Sprintf("failed to query database for songs:\r\n\r\n %v", err))
	}

	var songMap []fiber.Map

	for _, s := range songs {
		score, err := database.Queries.GetSongVoteScore(context.Background(), s.ID)
		if err != nil {
			return c.SendString(fmt.Sprintf("failed to query database for song %v score:\r\n\r\n %v", s.ID, err))
		}

		commentCount, err := database.Queries.CountCommentsBySong(context.Background(), s.ID)
		if err != nil {
			return c.SendString(fmt.Sprintf("failed to query database for song %v comment count:\r\n\r\n %v", s.ID, err))
		}

		songMap = append(songMap, fiber.Map{
			"link":          s.Link,
			"title":         "n/a tbd",
			"comment":       s.Comment.String,
			"created_at":    s.CreatedAt,
			"score":         score.Float64,
			"comment_count": commentCount,
		})
	}

	c.Render("list", fiber.Map{
		"title":       list.Name,
		"description": list.Description.String,
		"songs":       songMap,
	})

	// c.Render("list", fiber.Map{
	// 	"title":       "Demo List",
	// 	"description": "A really cool demo playlist!",
	// 	"songs": []fiber.Map{
	// 		{
	// 			"link":          "https://www.youtube.com/watch?v=3nlSDxvt6JU",
	// 			"title":         "Snail's House - Pixel Galaxy (Official MV)",
	// 			"comment":       "A really cool song I liked!",
	// 			"created_at":    "11:09 PM - 1 Jan 2016",
	// 			"score":         7,
	// 			"comment_count": 1,
	// 		},
	// 		{
	// 			"link":          "https://www.youtube.com/watch?v=uZf8a_tE00Q",
	// 			"title":         "Pixel Fantasy - Pixel Galaxy (Official MV)",
	// 			"comment":       "Another catchy song!",
	// 			"created_at":    "12:09 PM - 2 Jan 2016",
	// 			"score":         8,
	// 			"comment_count": 2,
	// 		},
	// 		{
	// 			"link":          "https://www.youtube.com/watch?v=n4U20g40jZc",
	// 			"title":         "Supercell - Snowflakes (Official MV)",
	// 			"comment":       "A beautiful song!",
	// 			"created_at":    "1:09 PM - 3 Jan 2016",
	// 			"score":         -2,
	// 			"comment_count": 3,
	// 		},
	// 		{
	// 			"link":          "https://www.youtube.com/watch?v=a_uG_h0v34o",
	// 			"title":         "Pixel Heart - Pixel Galaxy (Official MV)",
	// 			"comment":       "A touching song!",
	// 			"created_at":    "2:09 PM - 4 Jan 2016",
	// 			"score":         10,
	// 			"comment_count": 4,
	// 		},
	// 	},
	// })

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
