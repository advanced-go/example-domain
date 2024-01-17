package slo

import (
	"context"
	"errors"
	"github.com/advanced-go/core/runtime"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	postEntryHandlerLoc = PkgPath + ":postEntryHandler"
	createEntriesLoc    = PkgPath + ":createEntries"
	postRouteName       = "post-entry"
	postEntryLoc        = PkgPath + ":PostEntry"
)

func postEntryHandler[E runtime.ErrorHandler](ctx context.Context, h http.Header, method string, _ url.Values, body any) (t any, status runtime.Status) {
	var e E

	switch strings.ToUpper(method) {
	case http.MethodPut:
		var entries []EntryV1
		entries, status = createEntries(body)
		if !status.OK() {
			e.Handle(status, runtime.RequestId(h), postEntryHandlerLoc)
			return nil, status
		}
		if len(entries) == 0 {
			status = runtime.NewStatusError(runtime.StatusInvalidContent, postEntryHandlerLoc, errors.New("error: no entries found"))
			e.Handle(status, runtime.RequestId(h), postEntryHandlerLoc)
			return nil, status
		}
		status = addEntries(ctx, entries)
		if !status.OK() {
			e.Handle(status, runtime.RequestId(h), postEntryHandlerLoc)
		}
		return nil, status
	case http.MethodDelete:
		status = deleteEntries(ctx)
		if !status.OK() {
			e.Handle(status, runtime.RequestId(h), postEntryHandlerLoc)
		}
		return nil, status
	default:
		return nil, runtime.NewStatus(http.StatusMethodNotAllowed)
	}
}

func createEntries(body any) (entries []EntryV1, status runtime.Status) {
	if body == nil {
		return nil, runtime.NewStatus(runtime.StatusInvalidContent).AddLocation(createEntriesLoc)
	}

	switch ptr := body.(type) {
	case []EntryV1:
		entries = ptr
	case []byte:
		entries, status = runtime.New[[]EntryV1](ptr, nil)
		if !status.OK() {
			return nil, status.AddLocation(createEntriesLoc)
		}
	case *http.Request:
		entries, status = runtime.New[[]EntryV1](ptr, nil)
		if !status.OK() {
			return nil, status.AddLocation(createEntriesLoc)
		}
	case io.ReadCloser:
		entries, status = runtime.New[[]EntryV1](ptr, nil)
		if !status.OK() {
			return nil, status.AddLocation(createEntriesLoc)
		}
	default:
		return nil, runtime.NewStatusError(runtime.StatusInvalidContent, createEntriesLoc, runtime.NewInvalidBodyTypeError(body))
	}
	return entries, runtime.StatusOK()
}
