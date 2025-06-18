package handlers

import (
	"fmt"
	"sorn/internal/api"

	"github.com/gofiber/fiber/v2"
)


func AddTrackToDownload(c *fiber.Ctx) error {
	id := c.FormValue("id")
	fmt.Println(id)
	if id == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	err := api.QueueAlbum(id)
	if err != nil {
		fmt.Println(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusOK)
}
