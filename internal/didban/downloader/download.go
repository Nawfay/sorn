package downloader

import (
	// "sorn/internal/utils"
	// "sorn/internal/config"
	"fmt"
	"sorn/internal/db"
	"sorn/internal/didban/Logic"

	"github.com/kkdai/youtube/v2"
)

var YTClient *youtube.Client
var TmpPath string

func Init(arl string, tmpPath string) {
	SetARLCookie(arl)
	YTClient = &youtube.Client{}
	TmpPath = tmpPath
}


func DownloadTracks(item db.QueueItem) error {
	if item.Youtube && item.DeezerID != "" {
		track, err := logic.FetchTrack(item.DeezerID)
		if err != nil {
			return err
		}

		youtubeId, err := logic.DeezerToYtResolver(item.DeezerID)
		if err != nil {
			return err
		}
		
		finished, err := downloadTrackYt(YTClient, youtubeId, TmpPath, item.Path, track)
		if err != nil {
			return err
		}
		if !finished {
			return fmt.Errorf("track %s not found on YouTube", item.Title)
		}

		return nil
	}

	track, err := logic.FetchTrack(item.DeezerID)
	if err != nil {
		return err
	}

	finished, err := downloadTrackDeezer(track, TmpPath, item.Path)
	if err != nil {
		return err
	}
	if !finished {
		return fmt.Errorf("track %s not found on Deezer", item.Title)
	}

	return nil
}

