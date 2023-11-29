package entryv2

import (
	"net/url"
	"time"
)

type Entry struct {
	CreatedTS time.Time `json:"created-ts"`
	Traffic   string    `json:"traffic"`
	Start     time.Time `json:"start-time"`
	Duration  int       `json:"duration-ms"`

	RequestId string `json:"request-id"`

	// Request attributes
	Uri            string `json:"uri"` // {scheme}://{host}/{path} No query
	Protocol       string `json:"protocol"`
	Host           string `json:"host"`
	Path           string `json:"path"`
	Method         string `json:"method"`
	StatusCode     int32  `json:"status-code"`
	ThresholdFlags string `json:"threshold-flags"`
	Threshold      int    `json:"threshold"`
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
