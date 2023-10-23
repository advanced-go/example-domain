package activity

import (
	"encoding/json"
	"time"
)

type Entry struct {
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
type GetConstraints interface {
	[]Entry | []byte
}

var list []Entry

func GetEntries() []Entry {
	return list
}

func GetEntriesByType[T GetConstraints](act string) (T, error) {
	var l []Entry
	var t T
	var err error

	for _, v := range list {
		if v.ActivityType == act {
			l = append(l, v)
		}
	}
	switch ptr := any(&t).(type) {
	case *[]Entry:
		*ptr = l
	case *[]byte:
		buf, err2 := json.Marshal(l)
		if err2 == nil {
			*ptr = buf
		} else {
			err = err2
		}
	}
	return t, err
}

func AddEntry(e Entry) {
	list = append(list, e)
}

func deleteEntries() {
	list = []Entry{}
}
