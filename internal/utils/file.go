package utils

import (
	"fmt"
	"net/http"
	"os"

	"github.com/bogem/id3v2/v2"
	ffmpeg_go "github.com/u2takey/ffmpeg-go"
)


func GeneratePath(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}


func FetchCover(url string, outputPath string) error {
	// fmt.Println("Downloading image:", url, outputPath)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.ReadFrom(resp.Body)
	return err
}

func TagMP3(mp3Path, coverPath, title, artist, album, year string) error {
    // Open the MP3 file (read/write)
    tag, err := id3v2.Open(mp3Path, id3v2.Options{Parse: true})
    if err != nil {
        return fmt.Errorf("error opening mp3: %w", err)
    }
    defer tag.Close()

    // Set text metadata
    tag.SetTitle(title)
    tag.SetArtist(artist)
    tag.SetAlbum(album)
    tag.AddTextFrame(tag.CommonID("Year"), tag.DefaultEncoding(), year)

    // Load cover image from disk
    picBytes, err := os.ReadFile(coverPath)
    if err != nil {
        return fmt.Errorf("error reading cover image: %w", err)
    }

    // Create and set picture frame
    pic := id3v2.PictureFrame{
        Encoding:    id3v2.EncodingISO, // ISO-8859-1 is standard for images
        MimeType:    "image/jpeg",       // or "image/png"
        PictureType: id3v2.PTFrontCover,
        Description: "Cover",
        Picture:     picBytes,
    }
    tag.AddAttachedPicture(pic)

    // Save changes
    if err = tag.Save(); err != nil {
        return fmt.Errorf("error saving tag: %w", err)
    }
    return nil
}

func ConvertToMP4(input, output string, duration string) error {
	return ffmpeg_go.
		Input(input).
		Output(output, ffmpeg_go.KwArgs{
			// "c":        "copy",
            "t": duration,
			"movflags": "+faststart", // for web compatibility
		}).
		OverWriteOutput().
		Run()
}