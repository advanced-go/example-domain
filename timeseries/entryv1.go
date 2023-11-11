package timeseries

import (
	"net/url"
	"time"
)

var (
	ConrollerName = "ctrl"
)

type EntryV1 struct {
	CreatedTS time.Time
	Traffic   string
	Start     time.Time
	Duration  int

	// Use for ecosystem triage, not application triage.
	RequestId string

	// Request attributes
	Url         string // {scheme}://{host}/{path} No query
	Protocol    string // From timeseries
	Host        string // From timeseries
	Path        string // From timeseries
	Method      string
	StatusCode  int32
	StatusFlags string

	// Needed to verify client controller configuration matches configuration in cloud
	// Can this be replaced with a periodic audit?
	Timeout   int32
	RateLimit float64
	RateBurst int32
}

var listV1 []EntryV1

func getEntriesV1() []EntryV1 {
	return listV1
}

func addEntryV1(e []EntryV1) {
	for _, item := range e {
		listV1 = append(listV1, item)
	}
}

func deleteEntriesV1() {
	listV1 = []EntryV1{}
}

func queryEntriesV1(u *url.URL) []EntryV1 {
	return getEntriesV1()
}
