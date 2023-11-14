package slo

import (
	"github.com/google/uuid"
	"net/url"
	"time"
)

const (
	ControllerName = "ctrl"
)

type EntryV1 struct {
	CreatedTS time.Time
	Id        string
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

var list []EntryV1

func getEntries() []EntryV1 {
	return list
}

func getEntriesByController(ctrl string) []EntryV1 {
	for i := len(list) - 1; i >= 0; i-- {
		if list[i].Controller == ctrl {
			return []EntryV1{list[i]}
		}
	}
	return nil
}

func patchEntry(e EntryV1) {
	for i, _ := range list {
		if list[i].Controller == e.Controller {
			list[i] = e
			return
		}
	}
}

func addEntry(e []EntryV1) {
	for _, item := range e {
		if len(item.Id) == 0 {
			s, _ := uuid.NewUUID()
			item.Id = s.String()
		}
		//item.CreatedTS = time.Now().UTC()
		list = append(list, item)
	}
}

func deleteEntries() {
	list = []EntryV1{}
}

func queryEntries(u *url.URL) []EntryV1 {
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
