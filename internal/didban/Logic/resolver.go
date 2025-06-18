package logic

import (
	"fmt"

	"sorn/internal/utils"
)

func DeezerToYtResolver(deezerId string) (string, error) {

	track, err := FetchTrack(deezerId)
	if err != nil {
		return "", err
	}

	ytquery := fmt.Sprintf(track.Title + " " + track.Artist.Name + " lyrics")
	ytquery = utils.NormalizeStringForYT(ytquery)

	ytString, err := SearchYouTube(ytquery, 1)
	if err != nil {
		return "", err
	}

	return ytString[0], nil
}