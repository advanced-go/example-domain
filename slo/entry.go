package slo

import (
	"context"
	"github.com/advanced-go/core/io2"
	"github.com/advanced-go/core/json2"
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
	if location, ok := runtime.FileUrlFromContext(ctx); ok {
		return readEntry(location)
	}
	return list, runtime.StatusOK()
}

func getEntriesByController(ctx context.Context, ctrl string) ([]Entry, runtime.Status) {
	for i := len(list) - 1; i >= 0; i-- {
		if list[i].Controller == ctrl {
			return []Entry{list[i]}, runtime.StatusOK()
		}
	}
	return nil, runtime.StatusOK()
}

func addEntry(ctx context.Context, e []Entry) runtime.Status {
	if _, ok := runtime.FileUrlFromContext(ctx); ok {
		// Return OK, as we cannot go out of process
		return runtime.StatusOK()
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
	if _, ok := runtime.FileUrlFromContext(ctx); ok {
		// Return OK, as we cannot go out of process
		return runtime.StatusOK()
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

func readEntry(location string) (t []Entry, status runtime.Status) {
	buf, status2 := io2.ReadFileFromPath(location)
	if !status2.OK() {
		return t, status2.AddLocation(readEntryLoc)
	}
	status = json2.Unmarshal(buf, &t)
	if !status.OK() {
		return t, status.AddLocation(readEntryLoc)
	}
	return t, runtime.StatusOK()
}
