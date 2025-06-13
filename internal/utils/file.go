package utils


import (
	"net/http"
	"os"
)


func GeneratePath(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}


func DownloadImage(url string, outputPath string) error {

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