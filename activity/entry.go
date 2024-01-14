package activity

import (
	"context"
	"fmt"
	"github.com/advanced-go/core/runtime"
	"net/http"
	"net/url"
)

const (
	Type = "type"
)

type EntryV1 struct {
	//CreatedTS    time.Time
	ActivityID   string // Some form of UUID
	ActivityType string // trace|action
	Agent        string
	AgentUri     string // {host}:{agent}

	Assignment  string
	Controller  string
	Behavior    string
	Description string
}

var list []EntryV1

func getEntries(ctx context.Context) (t []EntryV1, status runtime.Status) {
	if url1, ok := lookup.Value("getEntries"); ok {
		return runtime.New[[]EntryV1](url1)
	}
	if len(list) == 0 {
		return list, runtime.NewStatus(http.StatusNotFound)
	}
	return list, runtime.StatusOK()
}

func getEntriesByType(ctx context.Context, act string) (t []EntryV1, status runtime.Status) {
	var l []EntryV1
	if url1, ok := lookup.Value("getEntriesByType"); ok {
		return runtime.New[[]EntryV1](url1)
	}
	for _, v := range list {
		if act == "" {
			l = append(l, v)
			continue
		}
		if v.ActivityType == act {
			l = append(l, v)
		}
	}
	if len(l) == 0 {
		return l, runtime.NewStatus(http.StatusNotFound)
	}
	return l, runtime.StatusOK()
}

func addEntries(ctx context.Context, e []EntryV1) runtime.Status {
	var status runtime.Status

	if url1, ok := lookup.Value("addEntries"); ok {
		return runtime.NewStatusFrom(url1)
	}
	for _, item := range e {
		//item.CreatedTS = time.Now().UTC()
		list = append(list, item)
		status = logActivity(ctx, item)
	}
	return status
}

func deleteEntries(ctx context.Context) runtime.Status {
	if url1, ok := lookup.Value("deleteEntries"); ok {
		return runtime.NewStatusFrom(url1)
	}
	list = []EntryV1{}
	return runtime.StatusOK()
}

func queryEntries(ctx context.Context, values url.Values) ([]EntryV1, runtime.Status) {
	var result []EntryV1
	var status runtime.Status

	name := ""
	if values != nil {
		name = values.Get(Type)
	}
	if len(name) != 0 {
		result, status = getEntriesByType(ctx, name)
	} else {
		result, status = getEntries(ctx)
	}
	return result, status
}

func logActivity(ctx context.Context, e EntryV1) runtime.Status {
	if url1, ok := lookup.Value("logActivity"); ok {
		return runtime.NewStatusFrom(url1)
	}
	s := fmt.Sprintf("{ \"activity\": \"%v\" \"agent\": \"%v\"  \"controller\": \"%v\"  \"message\": \"%v\"  }\n", e.ActivityType, e.Agent, e.Controller, e.Description)
	fmt.Printf("%v", s)
	return runtime.StatusOK()
}
