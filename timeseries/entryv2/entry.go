package entryv2

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
	Duration  int       `json:"duration-ms"`

	RequestId string `json:"request-id"`

	// Request attributes
	Uri            string `json:"uri"` // {scheme}://{host}/{path} No query
	Protocol       string `json:"protocol"`
	Host           string `json:"host"`
	Path           string `json:"path"`
	Method         string `json:"method"`
	StatusCode     int32  `json:"status-code"`
	ThresholdFlags string `json:"threshold-flags"`
	Threshold      int    `json:"threshold"`
}

const (
	readEntryLoc = PkgPath + ":readEntry"
)

var list []Entry

func getEntries(ctx context.Context) ([]Entry, *runtime.Status) {
	if url1, ok := lookup.Value("getEntries"); ok {
		return io2.New[[]Entry](url1, nil)
	}
	return list, runtime.StatusOK()
}

func addEntry(ctx context.Context, e []Entry) *runtime.Status {
	if url1, ok := lookup.Value("addEntries"); ok {
		return io2.NewStatusFrom(url1)
	}
	for _, item := range e {
		list = append(list, item)
	}
	return runtime.StatusOK()
}

func deleteEntries(ctx context.Context) *runtime.Status {
	if url1, ok := lookup.Value("deleteEntries"); ok {
		return io2.NewStatusFrom(url1)
	}
	list = []Entry{}
	return runtime.StatusOK()
}

func queryEntries(ctx context.Context, _ url.Values) ([]Entry, *runtime.Status) {
	return getEntries(ctx)
}
