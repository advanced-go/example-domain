package entryv2

import (
	"net/url"
	"time"
)

type Entry struct {
	CreatedTS time.Time
	Traffic   string
	Start     time.Time
	Duration  int

	RequestId string

	// Request attributes
	Url            string // {scheme}://{host}/{path} No query
	Protocol       string
	Host           string
	Path           string
	Method         string
	StatusCode     int32
	ThresholdFlags string
	Threshold      int
}

var list []Entry

func getEntries() []Entry {
	return list
}

func addEntry(e []Entry) {
	for _, item := range e {
		list = append(list, item)
	}
}

func deleteEntries() {
	list = []Entry{}
}

func queryEntries(u *url.URL) []Entry {
	return getEntries()
}
