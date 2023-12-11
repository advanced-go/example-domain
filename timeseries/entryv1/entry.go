package entryv1

import (
	"context"
	"github.com/advanced-go/core/io2"
	"github.com/advanced-go/core/json2"
	"github.com/advanced-go/core/runtime"
	"net/url"
	"time"
)

type Entry struct {
	CreatedTS time.Time `json:"created-ts"`
	Traffic   string    `json:"traffic"`
	Start     time.Time `json:"start-time"`
	Duration  int

	RequestId string

	// Request attributes
	Url         string // {scheme}://{host}/{path} No query
	Protocol    string
	Host        string
	Path        string
	Method      string
	StatusCode  int32
	StatusFlags string

	Timeout   int32
	RateLimit float64
	RateBurst int32
}

const (
	readEntryLoc = PkgPath + ":readEntry"
)

var list []Entry

func getEntries(ctx context.Context) ([]Entry, runtime.Status) {
	if location, ok := runtime.FileUrlFromContext(ctx); ok {
		return readEntry(location)
	}
	return list, runtime.StatusOK()
}

func addEntry(ctx context.Context, e []Entry) runtime.Status {
	if _, ok := runtime.FileUrlFromContext(ctx); ok {
		// Return OK, as we cannot go out of process
		return runtime.StatusOK()
	}
	for _, item := range e {
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
	return getEntries(ctx)
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
