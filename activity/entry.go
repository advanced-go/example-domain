package activity

import (
	"fmt"
	"net/url"
)

const (
	Type = "type"
)

var list []EntryV1

func getEntries() []EntryV1 {
	return list
}

func getEntriesByType(act string) []EntryV1 {
	var l []EntryV1

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

func addEntry(e []EntryV1) {
	for _, item := range e {
		//item.CreatedTS = time.Now().UTC()
		list = append(list, item)
		logActivity(item)
	}
}

func addItems(e []EntryV1) {
	for _, item := range e {
		//	item.CreatedTS = time.Now().UTC()
		list = append(list, item)
		fmt.Printf("%v\n", item)
	}
}

func addEntryTimes(e []EntryV1) {
	for _, item := range e {
		list = append(list, item)
		//fmt.Printf("%v\n", item)
	}
}

func deleteEntries() {
	list = []EntryV1{}
}

func queryEntries(u *url.URL) []EntryV1 {
	var result []EntryV1

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

func logActivity(e EntryV1) {
	s := fmt.Sprintf("{ \"activity\": \"%v\" \"agent\": \"%v\"  \"controller\": \"%v\"  \"message\": \"%v\"  }\n", e.ActivityType, e.Agent, e.Controller, e.Description)
	fmt.Printf("%v", s)

}
