
package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"sorn/internal/models"
)


func FetchArtist(id string) (models.Artist, error) {


	url := DeezerApi + "artist/" + id 

	resp, err := http.Get(url)
	if err != nil {
		return models.Artist{}, err 
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.Artist{}, errors.New("failed to fetch albums: " + resp.Status)
	}

	var trackResponse models.ArtistRaw
	err = json.NewDecoder(resp.Body).Decode(&trackResponse)
	if err != nil {
		return models.Artist{}, err
	}

	artist := models.Artist{
		ID:       trackResponse.ID,
		Name:    trackResponse.Name,
		NbAlbums: trackResponse.NbAlbums,
		PictureBig: trackResponse.PictureBig,
		PictureMedium: trackResponse.PictureMedium,
		PictureSmall: trackResponse.PictureSmall,
		PictureXL: trackResponse.PictureXL,
	}

	return artist, nil

}