package utils

import "time"

func DateFor(dateString string) time.Time {
	result, _ := time.Parse("2006-1-2", dateString)
	return result
}
