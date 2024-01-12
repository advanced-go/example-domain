package entryv1

import (
	"context"
	"errors"
	"github.com/advanced-go/core/access"
	"github.com/advanced-go/core/runtime"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	postHandlerLoc   = PkgPath + ":postHandler"
	createEntriesLoc = PkgPath + ":createEntries"
	postRouteName    = "post"
)

func postHandler[E runtime.ErrorHandler](ctx context.Context, h http.Header, method string, _ url.Values, body any) (t any, status runtime.Status) {
	var e E
	defer access.LogDeferred(access.InternalTraffic, access.NewRequest(h, method, postHandlerLoc), postRouteName, "", -1, "", &status)()

	switch strings.ToUpper(method) {
	case http.MethodPut:
		var entries []entry
		entries, status = createEntries(body)
		if !status.OK() {
			e.Handle(status, runtime.RequestId(h), postHandlerLoc)
			return nil, status
		}
		if len(entries) == 0 {
			status = runtime.NewStatusError(runtime.StatusInvalidContent, postHandlerLoc, errors.New("error: no entries found"))
			e.Handle(status, runtime.RequestId(h), "")
			return nil, status
		}
		status = addEntries(ctx, entries)
		if !status.OK() {
			e.Handle(status, runtime.RequestId(h), postHandlerLoc)
		}
		return nil, status
	case http.MethodDelete:
		status = deleteEntries(ctx)
		if !status.OK() {
			e.Handle(status, runtime.RequestId(h), postHandlerLoc)
		}
		return nil, status
	default:
		return nil, runtime.NewStatus(http.StatusMethodNotAllowed)
	}
}

func createEntries(body any) (entries []entry, status runtime.Status) {
	if body == nil {
		return nil, runtime.NewStatus(runtime.StatusInvalidContent).AddLocation(createEntriesLoc)
	}

	switch ptr := body.(type) {
	case []entry:
		entries = ptr
	case []byte:
		entries, status = runtime.New[[]entry](ptr)
		if !status.OK() {
			return nil, status.AddLocation(createEntriesLoc)
		}
	case io.ReadCloser:
		entries, status = runtime.New[[]entry](ptr)
		if !status.OK() {
			return nil, status.AddLocation(createEntriesLoc)
		}
	default:
		return nil, runtime.NewStatusError(runtime.StatusInvalidContent, createEntriesLoc, runtime.NewInvalidBodyTypeError(body))
	}
	return entries, runtime.StatusOK()
}
