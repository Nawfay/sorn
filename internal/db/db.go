package db

import (
    "fmt"
    "log"

    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

var DB *gorm.DB

type Artist struct {
    gorm.Model
    Name    string
    Tracked bool
    Youtube bool
    Albums  []Album
}

type Album struct {
    gorm.Model
    Title    string
    ArtistID uint
    Artist   Artist
    Youtube  bool
	Path     string
    Songs    []Song
}

type Song struct {
    gorm.Model
    Title    string
    Duration uint
    AlbumID  uint
    Album    Album
    FilePath string
}

func Connect() {
    // For now using sqlite, you can change to other DB by changing the driver here
    dbPath := "database.db"

    db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    err = db.AutoMigrate(&Artist{}, &Album{}, &Song{})
    if err != nil {
        log.Fatal("Failed to migrate database models:", err)
    }

    DB = db
    fmt.Println("Database connected and migrated!")
}
