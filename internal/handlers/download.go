package handlers


import (

	"sorn/internal/api"
	"github.com/gofiber/fiber/v2"
)


func AddTrackToDownload(c *fiber.Ctx) error {
	url := c.FormValue("url")
	if url == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	err := api.QueueAlbum(url)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusOK)
}
