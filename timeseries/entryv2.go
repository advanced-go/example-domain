package timeseries

import (
	"net/url"
	"time"
)

type EntryV2 struct {
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
	Threshold   int
}

var listV2 []EntryV2

func getEntriesV2() []EntryV2 {
	return listV2
}

func addEntryV2(e []EntryV2) {
	for _, item := range e {
		listV2 = append(listV2, item)
	}
}

func deleteEntriesV2() {
	listV2 = []EntryV2{}
}

func queryEntriesV2(u *url.URL) []EntryV2 {
	return getEntriesV2()
}
