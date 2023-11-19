package activity

import (
	"context"
	"github.com/advanced-go/core/http2"
	"github.com/advanced-go/core/io2"
	"github.com/advanced-go/core/json2"
	"github.com/advanced-go/core/runtime"
	"io"
	"net/http"
	"strings"
)

const (
	ContentLocation = "Content-Location"
	postLoc         = PkgUri + "/postEntryHandler"
	putEntryLoc     = PkgUri + "/putEntry"
)

type postEntryHandlerFn func(r *http.Request, body any) (any, runtime.Status)

func postEntryHandler[E runtime.ErrorHandler](ctx context.Context, r *http.Request, body any) (any, runtime.Status) {
	if r == nil {
		return nil, runtime.NewStatus(http.StatusBadRequest)
	}
	var e E
	if runtime.IsDebugEnvironment() {
		status2 := runtime.StatusFromContext(ctx)
		if status2 != nil {
			return nil, status2
		}
		location := r.Header.Get(http2.ContentLocation)
		if strings.HasPrefix(location, "file://") {
			return nil, runtime.NewStatusOK()
		}
	}
	statusVar := validateVariant(r)
	if !statusVar.OK() {
		e.Handle(statusVar, runtime.RequestId(r), postLoc)
		return nil, statusVar
	}
	switch strings.ToUpper(r.Method) {
	case http.MethodPut:
		return nil, e.Handle(putEntry(r.Header.Get(ContentLocation), body), runtime.RequestId(r), postLoc)
	case http.MethodDelete:
		return nil, e.Handle(deleteEntry(r.Header.Get(ContentLocation)), runtime.RequestId(r), postLoc)
	default:
		return nil, runtime.NewStatus(http.StatusMethodNotAllowed)
	}
}

func putEntry(variant string, body any) runtime.Status {
	if body == nil {
		runtime.NewStatus(runtime.StatusInvalidContent)
	}
	switch variant {
	case EntryV1Variant:
		var entries []EntryV1

		switch ptr := body.(type) {
		case []EntryV1:
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
			return runtime.NewStatus(runtime.StatusInvalidContent)
		}
		addEntry(entries)
		return runtime.NewStatusOK()
	default:
		return runtime.NewStatus(runtime.StatusInvalidContent)
	}
}

func deleteEntry(variant string) runtime.Status {
	switch variant {
	case EntryV1Variant:
		deleteEntries()
		return runtime.NewStatusOK()
	default:
		return runtime.NewStatus(runtime.StatusInvalidContent)
	}
}
