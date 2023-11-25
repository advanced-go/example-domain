package slo

import (
	"github.com/google/uuid"
	"net/url"
)

const (
	ControllerName = "ctrl"
)

var list []Entry

func getEntries() []Entry {
	return list
}

func getEntriesByController(ctrl string) []Entry {
	for i := len(list) - 1; i >= 0; i-- {
		if list[i].Controller == ctrl {
			return []Entry{list[i]}
		}
	}
	return nil
}

func patchEntry(e Entry) {
	for i, _ := range list {
		if list[i].Controller == e.Controller {
			list[i] = e
			return
		}
	}
}

func addEntry(e []Entry) {
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
	list = []Entry{}
}

func queryEntries(u *url.URL) []Entry {
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
