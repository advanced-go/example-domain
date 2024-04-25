package entryv1

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

func getEntries(ctx context.Context) ([]Entry, *core.Status) {
	if url1, ok := lookup.Value("getEntries"); ok {
		return json.New[[]Entry](url1, nil)
	}
	return list, core.StatusOK()
}

func addEntries(ctx context.Context, e []Entry) *core.Status {
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
