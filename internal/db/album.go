package db

import (
	"errors"
	"fmt"
	"sorn/internal/config"
	"sorn/internal/utils"

	"gorm.io/gorm"
)

func GetOrCreateAlbum(db *gorm.DB, title string, id int, artistID uint) (*Album, error) {
	var album Album

	// Check if album exists (by title and artistID)
	err := db.Where("title = ? AND artist_id = ?", title, artistID).First(&album).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// Album not found, create a new one

		// We need the artist name to build the album path
		artist, err := GetArtistByID(db, artistID)
		if err != nil {
			return nil, fmt.Errorf("failed to get artist for album path: %w", err)
		}

		// Build the album path using your utils function
		// Assuming BuildAlbumPath(downloadBasePath, artistName, albumTitle) string
		path := utils.BuildAlbumPath(config.Cfg.DownloadPath, artist.Name, title)

		album = Album{
			Title:    title,
			DeezerID: id,
			ArtistID: artistID,
			Path:     path,
			Youtube:  false, // default value
		}

		if err := db.Create(&album).Error; err != nil {
			return nil, err
		}
	} else if err != nil {
		// Some other database error
		return nil, err
	}

	return &album, nil
}

func GetAlbumByID(db *gorm.DB, id uint) (*Album, error) {
	var album Album

	err := db.Preload("Artist").Preload("Songs").First(&album, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil // no album found with that ID
	} else if err != nil {
		return nil, err // some other error occurred
	}

	return &album, nil
}
