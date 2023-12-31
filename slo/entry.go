package slo

import (
	"context"
	"github.com/advanced-go/core/runtime"
	"github.com/advanced-go/example-domain/slo/types"
	"github.com/google/uuid"
	"net/url"
)

const (
	ControllerName = "ctrl"
	readEntryLoc   = PkgPath + ":readEntry"
)

var list []types.Entry

func getEntries(ctx context.Context) ([]types.Entry, runtime.Status) {
	if url1, ok := lookup.Value("getEntries"); ok {
		return runtime.New[[]types.Entry](url1)
	}
	return list, runtime.StatusOK()
}

func getEntriesByController(ctx context.Context, ctrl string) ([]types.Entry, runtime.Status) {
	if url1, ok := lookup.Value("getEntriesByController"); ok {
		return runtime.New[[]types.Entry](url1)
	}
	for i := len(list) - 1; i >= 0; i-- {
		if list[i].Controller == ctrl {
			return []types.Entry{list[i]}, runtime.StatusOK()
		}
	}
	return nil, runtime.StatusOK()
}

func addEntries(ctx context.Context, e []types.Entry) runtime.Status {
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
	list = []types.Entry{}
	return runtime.StatusOK()
}

func queryEntries(ctx context.Context, values url.Values) ([]types.Entry, runtime.Status) {
	var result []types.Entry
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
