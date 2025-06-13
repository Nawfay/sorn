package handlers

import (
	"sorn/internal/api"
	"github.com/gofiber/fiber/v2"

	
)

func GetArtist(c *fiber.Ctx) error {
	artistID := c.Params("artist_id")
	artist, err := api.FetchArtist(artistID)
	if err != nil {
		return err
	}
	return c.JSON(artist)
}