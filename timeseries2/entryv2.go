package timeseries2

import (
	"net/url"
)

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
