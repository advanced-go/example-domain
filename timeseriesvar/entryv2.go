package timeseriesvar

import (
	"net/url"
)

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
