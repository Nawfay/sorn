package db

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

type QueueItem struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time

	DeezerID string `gorm:"uniqueIndex"` // Unique to avoid duplicates
	Title    string
	Artist   string
	Album    string
	URL      string
	Path     string
	Youtube  bool
	Status   string `gorm:"default:pending"` // pending, downloading, completed, failed
}

type Artist struct {
	gorm.Model
	Name     string
	DeezerID int `gorm:"uniqueIndex"` // unique artist ID from Deezer
	Tracked  bool
	Youtube  bool
	Albums   []Album `gorm:"foreignKey:ArtistID"`
}

type Album struct {
	gorm.Model
	Title    string
	DeezerID int `gorm:"uniqueIndex"` // unique album ID from Deezer
	ArtistID uint
	Artist   Artist `gorm:"foreignKey:ArtistID"`
	Youtube  bool
	Path     string
	Songs    []Song `gorm:"foreignKey:AlbumID"`
}

type Song struct {
	gorm.Model
	Title    string
	DeezerID string `gorm:"uniqueIndex"` // unique song ID from Deezer
	Duration uint
	AlbumID  uint
	Album    Album `gorm:"foreignKey:AlbumID"`
	FilePath string
	Youtube  bool
}

func Connect() {
	// For now using sqlite, you can change to other DB by changing the driver here
	dbPath := "database.db"

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = db.AutoMigrate(&Artist{}, &Album{}, &Song{}, &QueueItem{})
	if err != nil {
		log.Fatal("Failed to migrate database models:", err)
	}

	DB = db
	fmt.Println("Database connected and migrated!")
}
