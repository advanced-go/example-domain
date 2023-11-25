package activity

import (
	"fmt"
	"net/url"
)

const (
	Type = "type"
)

var list []Entry

func getEntries() []Entry {
	return list
}

func getEntriesByType(act string) []Entry {
	var l []Entry

	for _, v := range list {
		if act == "" {
			l = append(l, v)
			continue
		}
		if v.ActivityType == act {
			l = append(l, v)
		}
	}
	return l
}

func addEntry(e []Entry) {
	for _, item := range e {
		//item.CreatedTS = time.Now().UTC()
		list = append(list, item)
		logActivity(item)
	}
}

func addItems(e []Entry) {
	for _, item := range e {
		//	item.CreatedTS = time.Now().UTC()
		list = append(list, item)
		fmt.Printf("%v\n", item)
	}
}

func addEntryTimes(e []Entry) {
	for _, item := range e {
		list = append(list, item)
		//fmt.Printf("%v\n", item)
	}
}

func deleteEntries() {
	list = []Entry{}
}

func queryEntries(u *url.URL) []Entry {
	var result []Entry

	name := ""
	if u.Query() != nil {
		name = u.Query().Get(Type)
	}
	if len(name) != 0 {
		result = getEntriesByType(name)
	} else {
		result = getEntries()
	}
	return result
}

func logActivity(e Entry) {
	s := fmt.Sprintf("{ \"activity\": \"%v\" \"agent\": \"%v\"  \"controller\": \"%v\"  \"message\": \"%v\"  }\n", e.ActivityType, e.Agent, e.Controller, e.Description)
	fmt.Printf("%v", s)

}
