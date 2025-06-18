package api

import (
	"fmt"
	"sorn/internal/db"
	"sorn/internal/utils"
	"sorn/internal/config"
)



func QueueAlbum(id string) error {

	// get the album from Deezer
	albumData, err := FetchAlbum(id)
	if err != nil {
		return fmt.Errorf("failed to get album from Deezer: %w", err)
	}

	artistIdStr := fmt.Sprint(albumData.ArtistID)

	artistData, err := FetchArtist(artistIdStr)
	if err != nil {
		return fmt.Errorf("failed to get artist from Deezer: %w", err)
	}

	// get or create artist
	artist, err := db.GetOrCreateArtist(db.DB, albumData.Artist, albumData.ArtistID, false, false)
	if err != nil {
		return fmt.Errorf("failed to get/create artist: %w", err)
	}
	
	// create artist path
	path := utils.BuildArtistPath(config.Cfg.DownloadPath, artist.Name)
	err = utils.GeneratePath(path)
	if err != nil {
		return fmt.Errorf("failed to create artist path: %w", err)
	}

	// generare artist photo
	photoPath := fmt.Sprintf("%s/%s.jpg", path, utils.NormalizeFilename(artist.Name))
	err = utils.FetchCover(artistData.PictureMedium, photoPath)
	if err != nil {
		return fmt.Errorf("failed to download artist photo: %w", err)
	}

	// get or create album
	album, err := db.GetOrCreateAlbum(db.DB, albumData.Title, albumData.ID, artist.ID)
	if err != nil {
		return fmt.Errorf("failed to get/create album: %w", err)
	}

	// create album path
	path = utils.BuildAlbumPath(config.Cfg.DownloadPath, artist.Name, album.Title)
	err = utils.GeneratePath(path)
	if err != nil {
		return fmt.Errorf("failed to create album path: %w", err)
	}


	// Build track paths and enqueue tracks
	for _, t := range albumData.Tracks {
		trackPath := fmt.Sprintf("%s/%s.mp3", album.Path, utils.NormalizeFilename(t.Title))

		item := &db.QueueItem{
			DeezerID: fmt.Sprint(t.ID),
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
