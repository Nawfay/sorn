package queue

import (
	"fmt"
	// "log"
	"time"

	"sorn/internal/db"

	"github.com/nawfay/didban"
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
				time.Sleep(5 * time.Second) // Runs again in 5 minutes
				continue
			}

			SetStatus("working")
			fmt.Println("Downloading:", item.Title, "by", item.Artist)

			db.DB.Model(&item).Update("status", "downloading")

			// Convert internal QueueItem to didban QueueItem
			didbanItem := models.QueueItem{
				DeezerID: item.DeezerID,
				Title:    item.Title,
				Artist:   item.Artist,
				Album:    item.Album,
				Path:     item.Path,
				Youtube:  item.Youtube,
				Status:   item.Status,
			}

			err := didban.DownloadTracks(didbanItem)

			if err != nil {
				fmt.Println("Failed to download:", err)
				db.DB.Model(&item).Update("status", "failed")
				continue
			}

			db.DB.Model(&item).Updates(map[string]interface{}{
				"status": "completed",
				// "path":   item.Path, // optional: set final path
			})

			time.Sleep(10 * time.Second) // small delay before next
		}
	}()
}
