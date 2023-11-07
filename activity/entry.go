package activity

import (
	"fmt"
	"net/url"
	"time"
)

const (
	Type = "type"
)

type EntryV1 struct {
	CreatedTS    time.Time
	ActivityID   string // Some form of UUID
	ActivityType string // trace|action
	Agent        string
	AgentUri     string // {host}:{agent}

	Assignment  string
	FrameUri    string // {host}:{frame-name}
	Controller  string
	Behavior    string
	Description string
}

// GetConstraints - interface defining constraints for the Get function
// This could also be a representation that facilitates querying. Things like
// 1. What time did this occur?
// 2. Did this involve a specific entity?
// 3. ...
//type GetConstraints interface {
//	[]Entry | []byte
//}

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
		list = append(list, item)
		//fmt.Printf("%v\n", item)
	}
}

func addItems(e []EntryV1) {
	for _, item := range e {
		list = append(list, item)
		fmt.Printf("%v\n", item)
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
