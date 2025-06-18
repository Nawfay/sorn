package downloader

import (
	"fmt"
	"io"
	"os"

	"sorn/internal/models"
	"sorn/internal/utils"
	"github.com/kkdai/youtube/v2"
)

func downloadTrackYt(client *youtube.Client, videoID string, tmpPath string, path string, track *models.DidbanTrack) (bool, error) {

	// Fetch video metadata
	video, err := client.GetVideo(videoID)
	if err != nil {
		return false, fmt.Errorf("failed to get video info: %w", err)
	}

	video.Formats.Type("mp4")
	// Find best audio-only stream (highest bitrate)
	formats := video.Formats.WithAudioChannels() // only get videos with audio
	stream, _, err := client.GetStream(video, &formats[0])

	if err != nil {
		return false, fmt.Errorf("failed to get audio stream: %w", err)
	}
	defer stream.Close() // Close the stream when done

	tmpFile := fmt.Sprintf("%s/%s.tmp_audio", tmpPath, videoID)

	out, err := os.Create(tmpFile)
	if err != nil {
		return false, fmt.Errorf("failed to create temp file: %w", err)
	}
	defer out.Close()

	_, err = io.Copy(out, stream)
	if err != nil {
		return false, fmt.Errorf("failed to write audio stream: %w", err)
	}

	trackPath := fmt.Sprintf("%s/%s.mp3", path, utils.GenerateTrackTitle(track))

	utils.ConvertToMP4(tmpFile, trackPath, fmt.Sprintf("%d", video.Duration))
	os.Remove(tmpFile)

	err = utils.TagTrackWithMetadata(tmpPath, trackPath, videoID, track)
	if err != nil {
		os.Remove(trackPath)
		return false, fmt.Errorf("failed to tag MP3: %w", err)
	}

	return true, nil
}

