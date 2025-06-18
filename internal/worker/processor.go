package queue

import (
	"fmt"
	// "log"
	"time"

	"sorn/internal/db"
	"sorn/internal/didban/downloader"
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
				time.Sleep(5* time.Second) // Runs again in 5 minutes
				continue
			}

			SetStatus("working")
			fmt.Println("Downloading:", item.Title, "by", item.Artist)

			db.DB.Model(&item).Update("status", "downloading")
			err := downloader.DownloadTracks(item)
			
			if err != nil {
				fmt.Println("Failed to download:", err)
				db.DB.Model(&item).Update("status", "failed")
				continue
			}

			db.DB.Model(&item).Updates(map[string]interface{}{
				"status": "completed",
				// "path":   item.Path, // optional: set final path
			})

			time.Sleep(2 * time.Minute) // small delay before next
		}
	}()
}
