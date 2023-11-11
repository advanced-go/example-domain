package timeseries

import (
	"net/url"
	"time"
)

var (
	ConrollerName = "ctrl"
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
	RoutePct  int32 // A -1 value is for a disabled controller, everything else is the percentage of
	// traffic routed to secondary

}

var listV1 []EntryV1

func getEntries() []EntryV1 {
	return listV1
}

func addEntry(e []EntryV1) {
	for _, item := range e {
		listV1 = append(listV1, item)
	}
}

func deleteEntries() {
	listV1 = []EntryV1{}
}

func queryEntries(u *url.URL) []EntryV1 {
	return getEntries()
}
