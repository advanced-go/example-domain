package activity

import (
	"context"
	"fmt"
	"github.com/advanced-go/core/io2"
	"github.com/advanced-go/core/json2"
	"github.com/advanced-go/core/runtime"
	"net/http"
	"net/url"
)

const (
	Type         = "type"
	readEntryLoc = PkgPath + ":readEntry"
)

var list []Entry

func getEntries(ctx context.Context) (t []Entry, status runtime.Status) {
	if location, ok := ContentLocationFromContext(ctx); ok {
		return readEntry(location)
	}
	if len(list) == 0 {
		return list, runtime.NewStatus(http.StatusNotFound)
	}
	return list, runtime.StatusOK()
}

func getEntriesByType(ctx context.Context, act string) (t []Entry, status runtime.Status) {
	var l []Entry
	if location, ok := ContentLocationFromContext(ctx); ok {
		return readEntry(location)
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

func addEntry(ctx context.Context, e []Entry) runtime.Status {
	var status runtime.Status

	if _, ok := ContentLocationFromContext(ctx); ok {
		// Return OK, as we cannot go out of process
		return runtime.StatusOK()
	}
	for _, item := range e {
		//item.CreatedTS = time.Now().UTC()
		list = append(list, item)
		status = logActivity(ctx, item)
	}
	return status
}

func deleteEntries(ctx context.Context) runtime.Status {
	if _, ok := ContentLocationFromContext(ctx); ok {
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
		name = u.Query().Get(Type)
	}
	if len(name) != 0 {
		result, status = getEntriesByType(ctx, name)
	} else {
		result, status = getEntries(ctx)
	}
	return result, status
}

func logActivity(ctx context.Context, e Entry) runtime.Status {
	if _, ok := ContentLocationFromContext(ctx); ok {
		// Return OK, as we cannot go out of process
		return runtime.StatusOK()
	}
	s := fmt.Sprintf("{ \"activity\": \"%v\" \"agent\": \"%v\"  \"controller\": \"%v\"  \"message\": \"%v\"  }\n", e.ActivityType, e.Agent, e.Controller, e.Description)
	fmt.Printf("%v", s)
	return runtime.StatusOK()
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
