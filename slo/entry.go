package slo

import (
	"context"
	"github.com/advanced-go/core/runtime"
	"github.com/google/uuid"
	"net/url"
	"time"
)

const (
	ControllerName = "ctrl"
	readEntryLoc   = PkgPath + ":readEntry"
)

type EntryV1 struct {
	CreatedTS time.Time
	Id        string
	// What does this apply to
	Controller string

	// Types of SLOs
	// percentage of traffic : 10% or 10
	// latency percentile: 99/500ms
	Threshold   string // Either percentage of traffic, or latency percentile
	StatusCodes string // For percentage
}

var list []EntryV1

func getEntries(ctx context.Context) ([]EntryV1, runtime.Status) {
	if url1, ok := lookup.Value("getEntries"); ok {
		return runtime.New[[]EntryV1](url1)
	}
	return list, runtime.StatusOK()
}

func getEntriesByController(ctx context.Context, ctrl string) ([]EntryV1, runtime.Status) {
	if url1, ok := lookup.Value("getEntriesByController"); ok {
		return runtime.New[[]EntryV1](url1)
	}
	for i := len(list) - 1; i >= 0; i-- {
		if list[i].Controller == ctrl {
			return []EntryV1{list[i]}, runtime.StatusOK()
		}
	}
	return nil, runtime.StatusOK()
}

func addEntries(ctx context.Context, e []EntryV1) runtime.Status {
	if url1, ok := lookup.Value("addEntries"); ok {
		return runtime.NewStatusFrom(url1)
	}
	for _, item := range e {
		if len(item.Id) == 0 {
			s, _ := uuid.NewUUID()
			item.Id = s.String()
		}
		//item.CreatedTS = time.Now().UTC()
		list = append(list, item)
	}
	return runtime.StatusOK()
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
		name = values.Get(ControllerName)
	}
	if len(name) != 0 {
		result, status = getEntriesByController(ctx, name)
	} else {
		result, status = getEntries(ctx)
	}
	return result, status
}
