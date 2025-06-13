package handlers

import (
	"sorn/internal/api"
	"github.com/gofiber/fiber/v2"

	
)

func GetAlbum(c *fiber.Ctx) error {
	albumID := c.Params("album_id")
	album, err := api.FetchAlbum(albumID)
	if err != nil {
		return err
	}
	return c.JSON(album)
}
