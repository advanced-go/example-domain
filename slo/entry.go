package slo

import (
	"net/url"
)

const (
	ControllerName = "ctrl"
)

type entry struct {
	// What does this apply to
	Controller string

	// Types of SLOs
	// availability : 99% and 99.999%
	// percentage of traffic : 10% or 10
	// latency percentile: 99/500ms
	Threshold   string // Either percentage of traffic, or latency percentile
	StatusCodes string // For percentage
}

type Update struct {
	Controller string
	Behavior   string // Maybe properties, action? Can be empty
	Action     string // SQL set statement
}

var list []entry

func getEntries() []entry {
	return list
}

func getEntriesByController(ctrl string) []entry {
	for i := len(list) - 1; i >= 0; i-- {
		if list[i].Controller == ctrl {
			return []entry{list[i]}
		}
	}
	return nil
}

func patchEntry(e entry) {
	for i, _ := range list {
		if list[i].Controller == e.Controller {
			list[i] = e
			return
		}
	}
}

func addEntry(e []entry) {
	for _, item := range e {
		list = append(list, item)
	}
}

func deleteEntries() {
	list = []entry{}
}

func queryEntries(u *url.URL) []entry {
	name := ""
	if u.Query() != nil {
		name = u.Query().Get(ControllerName)
	}
	if len(name) != 0 {
		return getEntriesByController(name)
	} else {
		return getEntries()
	}
	return nil
}
