package db

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

// GetOrCreateArtist tries to find an artist by name.
// If it doesn’t exist, it creates a new artist with the given data.
func GetOrCreateArtist(db *gorm.DB, name string, id int, tracked bool, youtube bool) (*Artist, error) {
	var artist Artist

	err := db.Where("name = ?", name).First(&artist).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// Artist not found, create a new one
		artist = Artist{
			Name:     name,
			DeezerID: id,
			Tracked:  tracked,
			Youtube:  youtube,
		}
		if err := db.Create(&artist).Error; err != nil {
			fmt.Println("something wrong here1")
			return nil, err
		}
		return &artist, nil
	} else if err != nil {
		fmt.Println("something wrong here2")
		// Some other DB error
		return nil, err
	}

	// Artist found
	return &artist, nil
}
func GetArtistByID(db *gorm.DB, id uint) (*Artist, error) {
	var artist Artist

	err := db.Preload("Albums").Preload("Albums.Songs").First(&artist, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil // no artist found with that ID
	} else if err != nil {
		return nil, err // some other error occurred
	}

	return &artist, nil
}
