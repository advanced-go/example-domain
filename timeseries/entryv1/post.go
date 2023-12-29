package entryv1

import (
	"context"
	"errors"
	"github.com/advanced-go/core/http2"
	"github.com/advanced-go/core/runtime"
	"github.com/advanced-go/example-domain/json2"
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
		status = addEntries(ctx, entries)
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

func createEntries(body any) ([]Entry, runtime.Status) {
	if body == nil {
		return nil, runtime.NewStatus(runtime.StatusInvalidContent).AddLocation(createEntriesLoc)
	}
	var entries []Entry

	switch ptr := body.(type) {
	case []Entry:
		entries = ptr
	case []byte:
		status := json2.Unmarshal(ptr, &entries)
		if !status.OK() {
			return nil, status.AddLocation(createEntriesLoc)
		}
	case io.ReadCloser:
		buf, status := http2.ReadAll(ptr)
		if !status.OK() {
			return nil, status.AddLocation(createEntriesLoc)
		}
		status = json2.Unmarshal(buf, &entries)
		if !status.OK() {
			return nil, status.AddLocation(createEntriesLoc)
		}
	default:
		return nil, runtime.NewStatusError(runtime.StatusInvalidContent, createEntriesLoc, runtime.NewInvalidBodyTypeError(body))
	}
	return entries, runtime.StatusOK()
}
