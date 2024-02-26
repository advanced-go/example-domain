package activity

import (
	"context"
	"errors"
	"github.com/advanced-go/core/io2"
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

func postEntryHandler[E runtime.ErrorHandler](ctx context.Context, h http.Header, method string, _ url.Values, body any) (t any, status *runtime.Status) {
	var e E

	switch strings.ToUpper(method) {
	case http.MethodPut:
		var entries []EntryV1
		entries, status = createEntries(body)
		if !status.OK() {
			e.Handle(status, runtime.RequestId(h))
			return nil, status
		}
		if len(entries) == 0 {
			status = runtime.NewStatusError(runtime.StatusInvalidContent, errors.New("error: no entries found"), nil)
			e.Handle(status, runtime.RequestId(h))
			return nil, status
		}
		status = addEntries(ctx, entries)
		if !status.OK() {
			e.Handle(status, runtime.RequestId(h))
		}
		return nil, status
	case http.MethodDelete:
		status = deleteEntries(ctx)
		if !status.OK() {
			e.Handle(status, runtime.RequestId(h))
		}
		return nil, status
	default:
		return nil, runtime.NewStatus(http.StatusMethodNotAllowed)
	}
}

func createEntries(body any) (entries []EntryV1, status *runtime.Status) {
	if body == nil {
		return nil, runtime.NewStatus(runtime.StatusInvalidContent).AddLocation()
	}
	switch ptr := body.(type) {
	case []EntryV1:
		entries = ptr
	case []byte:
		entries, status = io2.New[[]EntryV1](ptr, nil)
		if !status.OK() {
			return nil, status.AddLocation()
		}
	case *http.Request:
		entries, status = io2.New[[]EntryV1](ptr.Body, nil)
		if !status.OK() {
			return nil, status.AddLocation()
		}
	case io.ReadCloser:
		entries, status = io2.New[[]EntryV1](ptr, nil)
		if !status.OK() {
			return nil, status.AddLocation()
		}
	default:
		return nil, runtime.NewStatusError(runtime.StatusInvalidContent, runtime.NewInvalidBodyTypeError(body), nil)
	}
	return entries, runtime.StatusOK()
}
