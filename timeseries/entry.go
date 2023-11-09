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
	CreatedTS  time.Time
	Traffic    string
	Start      time.Time
	Duration   int //time.Duration
	Controller string

	// Do we need an Origin Uri ?? A concatenation of the following fields. Maybe a controller Uri??
	Region     string
	Zone       string
	SubZone    string
	Service    string
	InstanceId string

	// Use for ecosystem triage, not application triage.
	RequestId string

	// Request attributes
	Url         string // {scheme}://{host}/{path} No query
	Route       string // primary|secondary Need routing groups
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

var list []EntryV1

func getEntries() []EntryV1 {
	return list
}

func addEntry(e []EntryV1) {
	for _, item := range e {
		//item.CreatedTS = time.Now().UTC()
		list = append(list, item)
	}
}

func getEntriesByController(ctrl string) []EntryV1 {
	var e []EntryV1

	for i, _ := range list {
		if list[i].Controller == ctrl {
			e = append(e, list[i])
		}
	}
	return e
}

func deleteEntries() {
	list = []EntryV1{}
}

func queryEntries(u *url.URL) []EntryV1 {
	name := ""
	if u.Query() != nil {
		name = u.Query().Get(ConrollerName)
	}
	if len(name) != 0 {
		return getEntriesByController(name)
	} else {
		return getEntries()
	}
	return nil
}
