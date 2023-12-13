package entryv1

import (
	"context"
	"github.com/advanced-go/core/io2"
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
	if url := runtime.LookupFromContext(ctx, ""); len(url) > 0 {
		return io2.ReadState[[]Entry](url)
	}
	return list, runtime.StatusOK()
}

func addEntry(ctx context.Context, e []Entry) runtime.Status {
	if url := runtime.LookupFromContext(ctx, ""); len(url) > 0 {
		return io2.ReadStatus(url)
	}
	for _, item := range e {
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

func queryEntries(ctx context.Context, _ url.Values) ([]Entry, runtime.Status) {
	return getEntries(ctx)
}
