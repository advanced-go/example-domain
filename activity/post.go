package activity

import (
	"github.com/advanced-go/core/io2"
	"github.com/advanced-go/core/json2"
	"github.com/advanced-go/core/runtime"
	"io"
	"net/http"
	"strings"
)

const (
	ContentLocation = "Content-Location"
)

type postEntryHandlerFn func(r *http.Request, body any) (any, *runtime.Status)

var (
	postLoc     = PkgUri + "/postEntryHandler"
	putEntryLoc = PkgUri + "/putEntry"
)

func postEntryHandler[E runtime.ErrorHandler](proxy postEntryHandlerFn, r *http.Request, body any) (any, *runtime.Status) {
	if r == nil {
		return nil, runtime.NewStatus(http.StatusBadRequest)
	}
	var e E

	if proxy != nil {
		return proxy(r, body)
	}
	statusVar := validateVariant(r)
	if !statusVar.OK() {
		e.Handle(statusVar, runtime.RequestId(r), postLoc)
		return nil, statusVar
	}
	switch strings.ToUpper(r.Method) {
	case http.MethodPut:
		status := putEntry(r.Header.Get(ContentLocation), body)
		if !status.OK() {
			e.Handle(status, runtime.RequestId(r), postLoc)
			return nil, status
		}
	case http.MethodDelete:
		status := deleteEntry(r.Header.Get(ContentLocation))
		if !status.OK() {
			e.Handle(status, runtime.RequestId(r), postLoc)
		}
		deleteEntries()
		return nil, runtime.NewStatusOK()
	default:
	}
	return nil, runtime.NewStatus(http.StatusMethodNotAllowed)
}

func putEntry(variant string, body any) *runtime.Status {
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
	default:
		return runtime.NewStatus(runtime.StatusInvalidContent)
	}
	return runtime.NewStatusOK()
}

func deleteEntry(variant string) *runtime.Status {
	switch variant {
	case EntryV1Variant:
		deleteEntries()
	default:
		return runtime.NewStatus(runtime.StatusInvalidContent)
	}
	return runtime.NewStatusOK()
}
