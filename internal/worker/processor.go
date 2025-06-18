package queue

import (
	"fmt"
	"log"
	"time"

	"sorn/internal/db"

	"github.com/nawfay/didban/didban"
	"github.com/nawfay/didban/didban/models"
)


func StartWorker() {
	go func() {
		for {
			var item db.QueueItem
			tx := db.DB.
				Where("status = ?", "pending").
				Order("created_at").
				First(&item)

			if tx.Error != nil {
				SetStatus("idle")
				fmt.Println("Queue empty, sleeping...")
				time.Sleep(5 * time.Minute) // Runs again in 5 minutes
				continue
			}

			SetStatus("working")
			fmt.Printf("Downloading: %s by %s\n", item.Title, item.Artist)

			db.DB.Model(&item).Update("status", "downloading")
			err := didban.DownloadTracks(models.QueueItem{
				DeezerID: item.DeezerID,
				Title:    item.Title,
				Artist:   item.Artist,
				Album:    item.Album,
				URL:      item.URL,
				Path:     item.Path,
				Youtube:  item.Youtube,
			})
			if err != nil {
				log.Printf("Failed to download: %v", err)
				db.DB.Model(&item).Update("status", "failed")
				continue
			}

			db.DB.Model(&item).Updates(map[string]interface{}{
				"status": "completed",
				// "path":   item.Path, // optional: set final path
			})

			time.Sleep(2 * time.Second) // small delay before next
		}
	}()
}
