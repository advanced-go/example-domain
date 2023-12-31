package entryv2

import (
	"context"
	"errors"
	"github.com/advanced-go/core/runtime"
	"github.com/advanced-go/example-domain/timeseries/types"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	postHandlerLoc   = PkgPath + ":postHandler"
	createEntriesLoc = PkgPath + ":createEntries"
)

func postHandler[E runtime.ErrorHandler](ctx context.Context, h http.Header, method string, _ url.Values, body any) (any, runtime.Status) {
	var e E

	switch strings.ToUpper(method) {
	case http.MethodPut:
		entries, status := createEntries(body)
		if !status.OK() {
			e.Handle(status, runtime.RequestId(h), postHandlerLoc)
			return nil, status
		}
		if len(entries) == 0 {
			status = runtime.NewStatusError(runtime.StatusInvalidContent, postHandlerLoc, errors.New("error: no entries found"))
			e.Handle(status, runtime.RequestId(h), "")
			return nil, status
		}
		status = addEntry(ctx, entries)
		if !status.OK() {
			e.Handle(status, runtime.RequestId(h), postHandlerLoc)
		}
		return nil, status
	case http.MethodDelete:
		status := deleteEntries(ctx)
		if !status.OK() {
			e.Handle(status, runtime.RequestId(h), postHandlerLoc)
		}
		return nil, status
	default:
		return nil, runtime.NewStatus(http.StatusMethodNotAllowed)
	}
}

func createEntries(body any) (entries []types.EntryV2, status runtime.Status) {
	if body == nil {
		return nil, runtime.NewStatus(runtime.StatusInvalidContent).AddLocation(createEntriesLoc)
	}

	switch ptr := body.(type) {
	case []types.EntryV2:
		entries = ptr
	case []byte:
		entries, status = runtime.New[[]types.EntryV2](ptr)
		if !status.OK() {
			return nil, status.AddLocation(createEntriesLoc)
		}
	case io.ReadCloser:
		entries, status = runtime.New[[]types.EntryV2](ptr)
		if !status.OK() {
			return nil, status.AddLocation(createEntriesLoc)
		}
	default:
		return nil, runtime.NewStatusError(runtime.StatusInvalidContent, createEntriesLoc, runtime.NewInvalidBodyTypeError(body))
	}
	return entries, runtime.StatusOK()
}
