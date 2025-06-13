package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"sorn/internal/models"
)


func FetchAlbum(id string) (models.Album, error) {


	url := DeezerApi + "album/" + id 

	resp, err := http.Get(url)
	if err != nil {
		return models.Album{}, err 
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.Album{}, errors.New("failed to fetch albums: " + resp.Status)
	}

	var trackResponse models.AlbumRaw
	err = json.NewDecoder(resp.Body).Decode(&trackResponse)
	if err != nil {
		return models.Album{}, err
	}

	tracks := extractAlbumTracks(trackResponse)


	album := models.Album{
		ID:       trackResponse.ID,
		Title:    trackResponse.Title,
		Artist:   trackResponse.Artist.Name,
		ArtistID: trackResponse.Artist.ID,
		Cover:    trackResponse.Cover,
		CoverBig: trackResponse.CoverBig,
		NbTracks: trackResponse.NumberOfTracks,
		Tracks:   tracks,
	}

	return album, nil
}
