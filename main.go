package main

import (
	"sorn/internal/config"
	"sorn/internal/db"
	"sorn/internal/handlers"
	queue "sorn/internal/worker"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/nawfay/didban"
)

func main() {

	config.Load()
	db.Connect()

	// Initialize didban library
	err := didban.Init(config.Cfg.ARL, config.Cfg.TmpPath)
	if err != nil {
		panic("Failed to initialize didban: " + err.Error())
	}

	queue.StartWorker()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: true,
	}))

	app.Get("/api/search/:query", handlers.SearchTrack)
	app.Get("/api/album/:album_id", handlers.GetAlbum)
	app.Get("/api/artist/:artist_id", handlers.GetArtist)

	app.Post("/api/download", handlers.AddTrackToDownload)

	app.Listen(":8080")
}
