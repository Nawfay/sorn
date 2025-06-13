package main

import (
    "sorn/internal/handlers"
	"sorn/internal/config"
	"sorn/internal/db"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

)

func main() {

	config.Load()
	db.Connect()

    app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowCredentials: true,
	}))

    app.Get("/api/search/:query", handlers.SearchTrack)
	app.Get("/api/album/:album_id", handlers.GetAlbum)
	app.Get("/api/artist/:artist_id", handlers.GetArtist)

    app.Listen(":8080")
}
