package api

import (
	"fmt"
	"sorn/internal/db"
	"sorn/internal/utils"
)

func QueueAlbum(id string) error {

	// get the album from Deezer
	albumData, err := FetchAlbum(id)
	if err != nil {
		return fmt.Errorf("failed to get album from Deezer: %w", err)
	}

	// get or create artist
	artist, err := db.GetOrCreateArtist(db.DB, albumData.Artist, albumData.ArtistID, false, false)
	if err != nil {
		return fmt.Errorf("failed to get/create artist: %w", err)
	}

	// get or create album
	album, err := db.GetOrCreateAlbum(db.DB, albumData.Title, albumData.ID, artist.ID)
	if err != nil {
		return fmt.Errorf("failed to get/create album: %w", err)
	}

	// Build track paths and enqueue tracks
	for _, t := range albumData.Tracks {
		trackPath := fmt.Sprintf("%s/%s.mp3", album.Path, utils.NormalizeFilename(t.Title))

		item := &db.QueueItem{
			DeezerID: t.ID,
			Title:    t.Title,
			Artist:   artist.Name,
			Album:    album.Title,
			URL:      "nil",
			Path:     trackPath,
			Youtube:  false,
			Status:   "pending",
		}

		_, err := db.EnqueueTrack(db.DB, item)
		if err != nil {
			return fmt.Errorf("failed to enqueue track %s: %w", t.Title, err)
		}
	}

	return nil
}
