package entryv1

import (
	"context"
	"github.com/advanced-go/core/runtime"
	"github.com/advanced-go/example-domain/timeseries/types"
	"net/url"
)

/*
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
*/
const (
	readEntryLoc = PkgPath + ":readEntry"
)

var list []types.EntryV1

func getEntries(ctx context.Context) ([]types.EntryV1, runtime.Status) {
	if url1, ok := lookup.Value("getEntries"); ok {
		return runtime.New[[]types.EntryV1](url1)
	}
	return list, runtime.StatusOK()
}

func addEntries(ctx context.Context, e []types.EntryV1) runtime.Status {
	if url1, ok := lookup.Value("addEntries"); ok {
		return runtime.NewStatusFrom(url1)
	}
	for _, item := range e {
		list = append(list, item)
	}
	return runtime.StatusOK()
}

func deleteEntries(ctx context.Context) runtime.Status {
	if url1, ok := lookup.Value("deleteEntries"); ok {
		return runtime.NewStatusFrom(url1)
	}
	list = []types.EntryV1{}
	return runtime.StatusOK()
}

func queryEntries(ctx context.Context, _ url.Values) ([]types.EntryV1, runtime.Status) {
	return getEntries(ctx)
}
