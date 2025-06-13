package db

import (
	"errors"
	"gorm.io/gorm"
)

func EnqueueTrack(db *gorm.DB, item *QueueItem) (*QueueItem, error) {
		var existing QueueItem
		err := db.Where("deezer_id = ?", item.DeezerID).First(&existing).Error
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			// Already exists, return it
			return &existing, nil
		}

		if err := db.Create(item).Error; err != nil {
			return nil, err
		}
		return item, nil
}