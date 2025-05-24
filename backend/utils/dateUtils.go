package utils

import "time"

// GetCurrentDate returns the current Unix timestamp
func GetCurrentDate() int64 {
	return time.Now().Unix()
}
