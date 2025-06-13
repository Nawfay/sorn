package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sorn/internal/models"
)


const DeezerApi = "http://api.deezer.com/"




func SearchAlbums(name string, limit int) ([]models.Album, error) {
	// Construct the API URL
	url := DeezerApi + "search/album?limit=" + fmt.Sprint(limit) + "&q=" + name

	// Perform the HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check for non-200 HTTP responses
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch albums: " + resp.Status)
	}

	// Decode the response body into the AlbumResponse struct
	var albumResponse models.AlbumResponse
	err = json.NewDecoder(resp.Body).Decode(&albumResponse)
	if err != nil {
		return nil, err
	}


	// Return the albums from the response
	return extractAlbums(albumResponse), nil
}	

func extractAlbumTracks(album models.AlbumRaw) []models.Track {

	url := album.Tracklist

	resp, err := http.Get(url)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil
	}

	var trackResponse models.AlbumTrackResponse
	err = json.NewDecoder(resp.Body).Decode(&trackResponse)
	if err != nil {
		return nil
	}

	var tracks []models.Track

	for _, track := range trackResponse.Data {

		tracks = append(tracks, models.Track{
			ID:       track.ID,
			Title:    track.Title,
			Duration: track.Duration,
			Artist:   track.Artist.Name,
			Album:    album.Title,
			AlbumID:  album.ID,
			AlbumCover: album.Cover,
			
			})
	}

	return tracks
}

func extractAlbums(albumResponse models.AlbumResponse) []models.Album {

	var albums []models.Album

	for _, album := range albumResponse.Data {
		
		tracks := extractAlbumTracks(album)

		albums = append(albums, models.Album{
			ID:       album.ID,
			Title:    album.Title,
			Artist:   album.Artist.Name,
			ArtistID: album.Artist.ID,
			Cover:    album.Cover,
			CoverBig: album.CoverBig,
			NbTracks: album.NumberOfTracks,
			Tracks:   tracks,
		})
	}

	return albums
}

func SearchTracks(name string, limit int) ([]models.Track, error) {
	// Construct the API URL
	url := DeezerApi + "search?limit=" + fmt.Sprint(limit) + "&q=" + name

	// Perform the HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check for non-200 HTTP responses
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch tracks: " + resp.Status)
	}

	// Decode the response body into the TrackResponse struct
	var trackResponse models.TrackResponse
	err = json.NewDecoder(resp.Body).Decode(&trackResponse)
	if err != nil {
		return nil, err
	}

	// Return the tracks from the response
	return extractTracks(trackResponse), nil
}

func extractTracks(trackResponse models.TrackResponse) []models.Track {

	var tracks []models.Track

	for _, track := range trackResponse.Data {
		tracks = append(tracks, models.Track{
			ID:       track.ID,
			Title:    track.Title,
			Duration: track.Duration,
			Artist:   track.Artist.Name,
			Album:    track.Album.Title,
			AlbumID:  track.Album.ID,
			AlbumCover: track.Album.Cover,
		})
	}

	return tracks
}
