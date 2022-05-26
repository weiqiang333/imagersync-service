package configVar

import (
	"bytes"
	"time"
)

var ConfigFile string

// v1 image rsync
type ImageRsyncConfig struct {
	Source string `json:"source"`
	Target string `json:"target"`
}
type ImageRsyncData struct {
	RsyncStatus         string    `json:"rsync_status"`
	RsyncError          interface{}    `json:"rsync_error"`
	RsyncInfo           string    `json:"rsync_info"`
	StartTime           time.Time `json:"start_time"`
	EndTime             time.Time `json:"end_time"`
	DurationTimeSeconds float64   `json:"duration_time_seconds"`
}
type ImageRsyncRequestBody struct {
	RsyncConfig ImageRsyncConfig `json:"rsync_config"`
}
type ImageRsyncResponseBody struct {
	RsyncData   ImageRsyncData   `json:"rsync_data"`
	RsyncConfig ImageRsyncConfig `json:"rsync_config"`
	Username    string           `json:"username"`
}

type ProcessData bytes.Buffer
