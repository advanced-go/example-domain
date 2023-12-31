package entryv2

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
*/
const (
	readEntryLoc = PkgPath + ":readEntry"
)

var list []types.EntryV2

func getEntries(ctx context.Context) ([]types.EntryV2, runtime.Status) {
	if url1, ok := lookup.Value("getEntries"); ok {
		return runtime.New[[]types.EntryV2](url1)
	}
	return list, runtime.StatusOK()
}

func addEntry(ctx context.Context, e []types.EntryV2) runtime.Status {
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
	list = []types.EntryV2{}
	return runtime.StatusOK()
}

func queryEntries(ctx context.Context, _ url.Values) ([]types.EntryV2, runtime.Status) {
	return getEntries(ctx)
}
