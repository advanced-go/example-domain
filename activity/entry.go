package activity

import (
	"encoding/json"
)

type Entry struct {
	Agent      string
	Assignment string
	Action     string
}

var list []Entry

func GetEntries() ([]byte, error) {
	return json.Marshal(list)
}

func AddEntry(e Entry) {
	list = append(list, e)
}
