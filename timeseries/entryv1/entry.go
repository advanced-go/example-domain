package entryv1

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
	Url         string // {scheme}://{host}/{path} No query
	Protocol    string
	Host        string
	Path        string
	Method      string
	StatusCode  int32
	StatusFlags string

	Timeout   int32
	RateLimit float64
	RateBurst int32
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
