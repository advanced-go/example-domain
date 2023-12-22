package activity

import (
	"context"
	"fmt"
	"github.com/advanced-go/core/io2"
	"github.com/advanced-go/core/runtime"
	"net/http"
	"net/url"
)

const (
	Type = "type"
)

var list []Entry

func getEntries(ctx context.Context) (t []Entry, status runtime.Status) {
	if url, ok := lookup.Value("getEntries"); ok {
		return io2.ReadValues[[]Entry](url)
	}
	if len(list) == 0 {
		return list, runtime.NewStatus(http.StatusNotFound)
	}
	return list, runtime.StatusOK()
}

func getEntriesByType(ctx context.Context, act string) (t []Entry, status runtime.Status) {
	var l []Entry
	if url, ok := lookup.Value("getEntriesByType"); ok {
		return io2.ReadValues[[]Entry](url)
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

func addEntries(ctx context.Context, e []Entry) runtime.Status {
	var status runtime.Status

	if url, ok := lookup.Value("addEntries"); ok {
		return io2.ReadStatus(url)
	}
	for _, item := range e {
		//item.CreatedTS = time.Now().UTC()
		list = append(list, item)
		status = logActivity(ctx, item)
	}
	return status
}

func deleteEntries(ctx context.Context) runtime.Status {
	if url, ok := lookup.Value("deleteEntries"); ok {
		return io2.ReadStatus(url)
	}
	list = []Entry{}
	return runtime.StatusOK()
}

func queryEntries(ctx context.Context, values url.Values) ([]Entry, runtime.Status) {
	var result []Entry
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

func logActivity(ctx context.Context, e Entry) runtime.Status {
	if url, ok := lookup.Value("logActivity"); ok {
		return io2.ReadStatus(url)
	}
	s := fmt.Sprintf("{ \"activity\": \"%v\" \"agent\": \"%v\"  \"controller\": \"%v\"  \"message\": \"%v\"  }\n", e.ActivityType, e.Agent, e.Controller, e.Description)
	fmt.Printf("%v", s)
	return runtime.StatusOK()
}
