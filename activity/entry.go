package activity

import (
	"encoding/json"
	"time"
)

type Entry struct {
	CreatedTS  time.Time
	ActivityID string // Some form of UUID
	Agent      string
	Assignment string
	FrameUri   string // {host}:{frame-name}
	Action     string
}

var list []Entry

func GetEntries() ([]byte, error) {
	return json.Marshal(list)
}

func AddEntry(e Entry) {
	list = append(list, e)
}
