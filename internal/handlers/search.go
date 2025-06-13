package handlers

import (
	"fmt"
	"sorn/internal/api"
	"sorn/internal/models"

	"github.com/gofiber/fiber/v2"

	
)

func SearchTrack(c *fiber.Ctx) error {

	name := c.Params("query")

	fmt.Println(name)

	tracks, err := api.SearchTracks(name, 4)
	if err != nil {
		return err
	}

	albums, err := api.SearchAlbums(name, 4)
	if err != nil {
		return err
	}

	if len(tracks) == 0 {
		tracks = []models.Track{}
	}

	if len(albums) == 0 {
		albums = []models.Album{}
	}

	searchResult := map[string]interface{}{
		"tracks": tracks,
		"albums": albums,
	}

	return c.JSON(searchResult)	
}
