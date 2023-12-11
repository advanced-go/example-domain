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
	if uri, ok := runtime.FileUrlFromContext(ctx); ok {
		return io2.ReadState[[]Entry](uri)
	}
	return list, runtime.StatusOK()
}

func getEntriesByController(ctx context.Context, ctrl string) ([]Entry, runtime.Status) {
	if uri, ok := runtime.FileUrlFromContext(ctx); ok {
		return io2.ReadState[[]Entry](uri)
	}
	for i := len(list) - 1; i >= 0; i-- {
		if list[i].Controller == ctrl {
			return []Entry{list[i]}, runtime.StatusOK()
		}
	}
	return nil, runtime.StatusOK()
}

func addEntry(ctx context.Context, e []Entry) runtime.Status {
	if uri, ok := runtime.FileUrlFromContext(ctx); ok {
		return io2.ReadStatus(uri)
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
	if uri, ok := runtime.FileUrlFromContext(ctx); ok {
		return io2.ReadStatus(uri)
	}
	list = []Entry{}
	return runtime.StatusOK()
}

func queryEntries(ctx context.Context, u *url.URL) ([]Entry, runtime.Status) {
	var result []Entry
	var status runtime.Status

	name := ""
	if u.Query() != nil {
		name = u.Query().Get(ControllerName)
	}
	if len(name) != 0 {
		result, status = getEntriesByController(ctx, name)
	} else {
		result, status = getEntries(ctx)
	}
	return result, status
}
