package activity

import (
	"time"
)

const (
	Type = "type"
)

type entry struct {
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

var list []entry

func getEntries() []entry {
	return list
}

func getEntriesByType(act string) []entry {
	var l []entry

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

func addEntry(e []entry) {
	for _, item := range e {
		list = append(list, item)
	}
}

func deleteEntries() {
	list = []entry{}
}
