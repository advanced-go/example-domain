package slo

import (
	"context"
	"github.com/advanced-go/core/io2"
	"github.com/advanced-go/core/runtime"
	"github.com/google/uuid"
	"net/url"
)

const (
	ControllerName = "ctrl"
	readEntryLoc   = PkgPath + ":readEntry"
)

var list []Entry

func getEntries(ctx context.Context) ([]Entry, runtime.Status) {
	if url := runtime.LookupFromContext(ctx, ""); len(url) > 0 {
		return io2.ReadState[[]Entry](url)
	}
	return list, runtime.StatusOK()
}

func getEntriesByController(ctx context.Context, ctrl string) ([]Entry, runtime.Status) {
	if url := runtime.LookupFromContext(ctx, ""); len(url) > 0 {
		return io2.ReadState[[]Entry](url)
	}
	for i := len(list) - 1; i >= 0; i-- {
		if list[i].Controller == ctrl {
			return []Entry{list[i]}, runtime.StatusOK()
		}
	}
	return nil, runtime.StatusOK()
}

func addEntry(ctx context.Context, e []Entry) runtime.Status {
	if url := runtime.LookupFromContext(ctx, ""); len(url) > 0 {
		return io2.ReadStatus(url)
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
	if url := runtime.LookupFromContext(ctx, ""); len(url) > 0 {
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
		name = values.Get(ControllerName)
	}
	if len(name) != 0 {
		result, status = getEntriesByController(ctx, name)
	} else {
		result, status = getEntries(ctx)
	}
	return result, status
}
