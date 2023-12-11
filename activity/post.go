package activity

import (
	"context"
	"github.com/advanced-go/core/io2"
	"github.com/advanced-go/core/json2"
	"github.com/advanced-go/core/runtime"
	"io"
	"net/http"
	"strings"
)

const (
	postEntryHandlerLoc = PkgPath + ":postEntryHandler"
	putEntryLoc         = PkgPath + ":putEntry"
)

func postEntryHandler[E runtime.ErrorHandler](r *http.Request, body any) (any, runtime.Status) {
	var e E

	if r == nil {
		return nil, runtime.NewStatus(http.StatusBadRequest)
	}
	ctx := runtime.NewFileUrlContext(nil, r.URL.String())
	switch strings.ToUpper(r.Method) {
	case http.MethodPut:
		status := putEntry(ctx, body)
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

func putEntry(ctx context.Context, body any) runtime.Status {
	if body == nil {
		runtime.NewStatus(runtime.StatusInvalidContent).AddLocation(putEntryLoc)
	}
	var entries []Entry

	switch ptr := body.(type) {
	case []Entry:
		entries = ptr
	case []byte:
		status := json2.Unmarshal(ptr, &entries)
		if !status.OK() {
			return status.AddLocation(putEntryLoc)
		}
	case io.ReadCloser:
		buf, status := io2.ReadAll(ptr)
		if !status.OK() {
			return status.AddLocation(putEntryLoc)
		}
		status = json2.Unmarshal(buf, &entries)
		if !status.OK() {
			return status.AddLocation(putEntryLoc)
		}
	default:
		return runtime.NewStatusError(runtime.StatusInvalidContent, putEntryLoc, runtime.NewInvalidBodyTypeError(body))
	}
	if len(entries) == 0 {
		return runtime.NewStatus(http.StatusNotFound)
	}
	return addEntry(ctx, entries)
}
