package main

import (
    "sorn/internal/handlers"
	"sorn/internal/config"
	"sorn/internal/db"
	"sorn/internal/didban/downloader"
	"sorn/internal/worker"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

)

func main() {

	config.Load()
	db.Connect()
	downloader.Init(config.Cfg.ARL, config.Cfg.TmpPath)

	queue.StartWorker()

    app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowCredentials: true,
	}))

    app.Get("/api/search/:query", handlers.SearchTrack)
	app.Get("/api/album/:album_id", handlers.GetAlbum)
	app.Get("/api/artist/:artist_id", handlers.GetArtist)

	app.Post("/api/download", handlers.AddTrackToDownload)

    app.Listen(":8080")
}
