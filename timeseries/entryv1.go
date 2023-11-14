package timeseries

import (
	"net/url"
)

var (
	ConrollerName = "ctrl"
)

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
