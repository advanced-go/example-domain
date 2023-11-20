package timeseries

import (
	"context"
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

const (
	ContentLocation = "Content-Location"
	postLoc         = PkgUri + "/postEntryHandler"
	putLoc          = PkgUri + "/putEntry"
)

func postEntryHandler(ctx context.Context, r *http.Request, body any) (any, runtime.Status) {
	if r == nil {
		return nil, runtime.NewStatus(runtime.StatusInvalidContent)
	}

	if runtime.IsDebugEnvironment() {
		status2 := runtime.StatusFromContext(ctx)
		if status2 != nil {
			return nil, status2.AddLocation(postLoc)
		}
		location := r.Header.Get(http2.ContentLocation)
		if strings.HasPrefix(location, "file://") {
			// Need to deserialize return any
			return nil, runtime.NewStatusOK()
		}
	}
	statusVar := validateVariant(r)
	if !statusVar.OK() {
		//e.Handle(statusVar, runtime.RequestId(r), postLoc)
		return nil, statusVar.AddLocation(postLoc)
	}
	switch strings.ToUpper(r.Method) {
	case http.MethodPut:
		//return nil, e.Handle(putEntry(r.Header.Get(http2.ContentLocation), body), runtime.RequestId(r), postLoc)
		return nil, putEntry(r.Header.Get(http2.ContentLocation), body).AddLocation(postLoc)
	case http.MethodDelete:
		//return nil, e.Handle(deleteEntry(r.Header.Get(http2.ContentLocation)), runtime.RequestId(r), postLoc)
		return nil, deleteEntry(r.Header.Get(http2.ContentLocation)).AddLocation(postLoc)
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
		return runtime.NewStatusOK()
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
		return runtime.NewStatusOK()
	default:
		return runtime.NewStatus(runtime.StatusInvalidContent)
	}
}

func deleteEntry(variant string) runtime.Status {
	switch variant {
	case EntryV1Variant:
		deleteEntriesV1()
		return runtime.NewStatusOK()
	case EntryV2Variant:
		deleteEntriesV2()
		return runtime.NewStatusOK()
	default:
		return runtime.NewStatus(runtime.StatusInvalidContent)
	}
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
