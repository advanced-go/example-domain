package activity

import (
	"context"
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/json"
	"net/http"
	"net/url"
)

const (
	Type = "type"
)

var list []EntryV1

func getEntries(ctx context.Context) (t []EntryV1, status *core.Status) {
	if url1, ok := lookup.Value("getEntries"); ok {
		return json.New[[]EntryV1](url1, nil)
	}
	if len(list) == 0 {
		return list, core.NewStatus(http.StatusNotFound)
	}
	return list, core.StatusOK()
}

func getEntriesByType(ctx context.Context, act string) (t []EntryV1, status *core.Status) {
	var l []EntryV1
	if url1, ok := lookup.Value("getEntriesByType"); ok {
		return json.New[[]EntryV1](url1, nil)
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
		return l, core.NewStatus(http.StatusNotFound)
	}
	return l, core.StatusOK()
}

func addEntries(ctx context.Context, e []EntryV1) *core.Status {
	var status *core.Status

	if url1, ok := lookup.Value("addEntries"); ok {
		return json.NewStatusFrom(url1)
	}
	for _, item := range e {
		//item.CreatedTS = time.Now().UTC()
		list = append(list, item)
		status = logActivity(ctx, item)
	}
	return status
}

func deleteEntries(ctx context.Context) *core.Status {
	if url1, ok := lookup.Value("deleteEntries"); ok {
		return json.NewStatusFrom(url1)
	}
	list = []EntryV1{}
	return core.StatusOK()
}

func queryEntries(ctx context.Context, values url.Values) ([]EntryV1, *core.Status) {
	var result []EntryV1
	var status *core.Status

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

func logActivity(ctx context.Context, e EntryV1) *core.Status {
	if url1, ok := lookup.Value("logActivity"); ok {
		return json.NewStatusFrom(url1)
	}
	s := fmt.Sprintf("{ \"activity\": \"%v\" \"agent\": \"%v\"  \"controller\": \"%v\"  \"message\": \"%v\"  }\n", e.ActivityType, e.Agent, e.Controller, e.Description)
	fmt.Printf("%v", s)
	return core.StatusOK()
}
