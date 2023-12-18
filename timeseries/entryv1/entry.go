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
	if urls, ok := lookup("getEntries"); ok {
		return io2.ReadResults[[]Entry](urls)
	}
	return list, runtime.StatusOK()
}

func addEntries(ctx context.Context, e []Entry) runtime.Status {
	if urls, ok := lookup("addEntries"); ok {
		return io2.ReadStatus(urls)
	}
	for _, item := range e {
		list = append(list, item)
	}
	return runtime.StatusOK()
}

func deleteEntries(ctx context.Context) runtime.Status {
	if urls, ok := lookup("deleteEntries"); ok {
		return io2.ReadStatus(urls)
	}
	list = []Entry{}
	return runtime.StatusOK()
}

func queryEntries(ctx context.Context, _ url.Values) ([]Entry, runtime.Status) {
	return getEntries(ctx)
}
