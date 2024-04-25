package entryv2

import (
	"context"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/json"
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

func getEntries(ctx context.Context) ([]Entry, *core.Status) {
	if url1, ok := lookup.Value("getEntries"); ok {
		return json.New[[]Entry](url1, nil)
	}
	return list, core.StatusOK()
}

func addEntry(ctx context.Context, e []Entry) *core.Status {
	if url1, ok := lookup.Value("addEntries"); ok {
		return json.NewStatusFrom(url1)
	}
	for _, item := range e {
		list = append(list, item)
	}
	return core.StatusOK()
}

func deleteEntries(ctx context.Context) *core.Status {
	if url1, ok := lookup.Value("deleteEntries"); ok {
		return json.NewStatusFrom(url1)
	}
	list = []Entry{}
	return core.StatusOK()
}

func queryEntries(ctx context.Context, _ url.Values) ([]Entry, *core.Status) {
	return getEntries(ctx)
}
