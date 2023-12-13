package activity

import (
	"errors"
	"github.com/advanced-go/core/io2"
	"github.com/advanced-go/core/json2"
	"github.com/advanced-go/core/runtime"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	postEntryHandlerLoc = PkgPath + ":postEntryHandler"
	createEntriesLoc    = PkgPath + ":createEntries"
)

func postEntryHandler[E runtime.ErrorHandler](r *http.Request, _ url.Values, body any) (any, runtime.Status) {
	var e E

	if r == nil {
		return nil, runtime.NewStatus(http.StatusBadRequest)
	}
	ctx := runtime.NewFileUrlContext(nil, r.URL.String())
	switch strings.ToUpper(r.Method) {
	case http.MethodPut:
		entries, status := createEntries(body)
		if !status.OK() {
			e.Handle(status, runtime.RequestId(r), postEntryHandlerLoc)
			return nil, status
		}
		if len(entries) == 0 {
			status = runtime.NewStatusError(runtime.StatusInvalidContent, postEntryHandlerLoc, errors.New("error: no entries found"))
			e.Handle(status, runtime.RequestId(r), postEntryHandlerLoc)
			return nil, status
		}
		status = addEntry(ctx, entries)
		if !status.OK() {
			e.Handle(status, runtime.RequestId(r), postEntryHandlerLoc)
		}
		return nil, status
	case http.MethodDelete:
		status := deleteEntries(ctx)
		if !status.OK() {
			e.Handle(status, runtime.RequestId(r), postEntryHandlerLoc)
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
		buf, status := io2.ReadAll(ptr)
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
