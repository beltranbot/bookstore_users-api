package dateutils

import "time"

const (
	apiDateLayout = "2006-01-02T15:04:05Z"
)

// GetNow returns current timestamp as a time.Time
func GetNow() time.Time {
	return time.Now().UTC()
}

// GetNowString returns current timestamp as a string
func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}
