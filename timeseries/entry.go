package timeseries

import (
	"encoding/json"
	"time"
)

// Research - https://www.faa.gov/nextgen

// Entry - timeseries struct provides support for the following:
//      1. Resiliency - Comparing real time metrics against controller attributes and making appropriate
//                      controller configuration changes.
//      2. Traffic Analytics - Provide real time diagnostics of a distributed ecosystem, and affect
//                                      appropriate changes to keep ecosystem resilient and highly available.
//                                      Can routing be changed to avoid failures?
//
// Notes:

type Entry struct {
	Traffic    string
	Start      time.Time
	Duration   time.Duration
	Controller string

	// Do we need an Origin Uri ?? A concatenation of the following fields. Maybe a controller Uri??
	Region     string
	Zone       string
	SubZone    string
	Service    string
	InstanceId string

	// Need this ??
	RequestId string

	// Request attributes
	Url         string // {scheme}://{host}/{path} No query
	Route       string // primary|secondary
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
	RoutePct  int32 // A -1 value is for a disabled controller, everything else is the percentage of
	// traffic routed to secondary

}

var list []Entry

func GetEntries() ([]byte, error) {
	return json.Marshal(list)
}

func AddEntry(e Entry) {
	list = append(list, e)
}

func GetEntriesByController(ctrl string) []*Entry {
	var e []*Entry

	for i, _ := range list {
		if list[i].Controller == ctrl {
			e = append(e, &list[i])
		}
	}
	return e
}

func deleteEntries() {
	list = []Entry{}
}
