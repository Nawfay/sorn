package queue

import (
	"sync/atomic"
)

var queueStatus atomic.Value

func init() {
	queueStatus.Store("idle") // Initial status
}

func SetStatus(status string) {
	queueStatus.Store(status)
}

func GetStatus() string {
	val := queueStatus.Load()
	if val == nil {
		return "unknown"
	}
	return val.(string)
}
