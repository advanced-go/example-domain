package timeseries

import (
	"github.com/advanced-go/core/http2"
	"github.com/advanced-go/core/io2"
	"github.com/advanced-go/core/json2"
	"github.com/advanced-go/core/runtime"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type pkg struct{}

type postEntryHandlerFn func(r *http.Request, body any) (any, *runtime.Status)

const (
	postLoc = PkgUri + "/postEntryHandler"
	putLoc  = PkgUri + "/putEntry"
)

func postEntryHandler[E runtime.ErrorHandler](proxy postEntryHandlerFn, r *http.Request, body any) (any, *runtime.Status) {
	if r == nil {
		return nil, runtime.NewStatus(runtime.StatusInvalidContent)
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
		status := putEntry(r.Header.Get(http2.ContentLocation), body)
		if !status.OK() {
			e.Handle(status, runtime.RequestId(r), postLoc)
		}
		return nil, status
	case http.MethodDelete:
		status := deleteEntry(r.Header.Get(http2.ContentLocation))
		if !status.OK() {
			e.Handle(status, runtime.RequestId(r), postLoc)
		}
		return nil, status
	default:
		return nil, runtime.NewStatus(http.StatusMethodNotAllowed)
	}
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
				return status.AddLocation(putLoc)
			}
		case io.ReadCloser:
			buf, status := io2.ReadAll(ptr)
			if !status.OK() {
				return status.AddLocation(putLoc)
			}
			status = json2.Unmarshal(buf, &entries)
			if !status.OK() {
				return status.AddLocation(putLoc)
			}
		default:
			return runtime.NewStatusError(runtime.StatusInvalidContent, putLoc, runtime.NewInvalidBodyTypeError(body))
		}
		if len(entries) == 0 {
			return runtime.NewStatus(runtime.StatusInvalidContent)
		}
		addEntryV1(entries)
	case EntryV2Variant:
		var entries []EntryV2

		switch ptr := body.(type) {
		case []EntryV2:
			entries = ptr
		case []byte:
			status := json2.Unmarshal(ptr, &entries)
			if !status.OK() {
				return status.AddLocation(putLoc)
			}
		case io.ReadCloser:
			buf, status := io2.ReadAll(ptr)
			if !status.OK() {
				return status.AddLocation(putLoc)
			}
			status = json2.Unmarshal(buf, &entries)
			if !status.OK() {
				return status.AddLocation(putLoc)
			}
		default:
			return runtime.NewStatusError(runtime.StatusInvalidContent, putLoc, runtime.NewInvalidBodyTypeError(body))
		}
		if len(entries) == 0 {
			return runtime.NewStatus(runtime.StatusInvalidContent)
		}
		addEntryV2(entries)
	default:
		return runtime.NewStatus(runtime.StatusInvalidContent)
	}
	return runtime.NewStatusOK()
}

func deleteEntry(variant string) *runtime.Status {
	switch variant {
	case EntryV1Variant:
		deleteEntriesV1()
	case EntryV2Variant:
		deleteEntriesV2()
	default:
		return runtime.NewStatus(runtime.StatusInvalidContent)
	}
	return runtime.NewStatusOK()
}

func verifyVariant(u *url.URL, variant string) string {
	if u != nil {
		v := u.Query().Get("v")
		if len(v) > 0 {
			if v == "1" {
				return EntryV1Variant
			}
			if v == "2" {
				return EntryV2Variant
			}
		}
	}
	if len(variant) == 0 {
		return EntryV1Variant
	}
	return variant
}
