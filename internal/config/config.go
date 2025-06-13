package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var Cfg *Config

type Config struct {
	DownloadPath string
}

func Load() {
	// Load variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found or couldn't be loaded (that's okay in production)")
	}

	Cfg = &Config{
		DownloadPath: os.Getenv("DOWNLOAD_PATH"),
	}

	if Cfg.DownloadPath == "" {
		log.Fatal("DOWNLOAD_PATH is not set")
	}
}
