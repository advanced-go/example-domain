package timeseries

import (
	"github.com/go-ai-agent/core/http2"
	"github.com/go-ai-agent/core/io2"
	"github.com/go-ai-agent/core/json2"
	"github.com/go-ai-agent/core/log2"
	"github.com/go-ai-agent/core/runtime"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type pkg struct{}

var (
	postWrapper = log2.WrapPost(newPostEntryHandler[runtime.LogError]())
	postLoc     = PkgUri + "/postEntryHandler"
	putLoc      = PkgUri + "/putEntry"
)

func newPostEntryHandler[E runtime.ErrorHandler]() runtime.PostHandler {
	return func(ctx any, r *http.Request, body any) (any, *runtime.Status) {
		return postEntryHandler[E](ctx, r, body)
	}
}

func postEntryHandler[E runtime.ErrorHandler](ctx any, r *http.Request, body any) (any, *runtime.Status) {
	if r == nil {
		return nil, runtime.NewStatus(runtime.StatusInvalidContent)
	}
	var e E

	statusVar := validateVariant(r)
	if !statusVar.OK() {
		e.Handle(statusVar, runtime.RequestId(r), httpLoc)
		return nil, statusVar
	}
	if runtime.IsDebugEnvironment() {
		if fn := http2.PostHandlerProxy(ctx); fn != nil {
			return fn(ctx, r, body)
		}
	}
	switch strings.ToUpper(r.Method) {
	case http.MethodPut:
		status := putEntry(nil, r.Header.Get(http2.ContentLocation), body)
		if !status.OK() {
			e.Handle(status, runtime.RequestId(ctx), postLoc)
			return nil, status
		}
		return nil, runtime.NewStatusOK()
	case http.MethodDelete:
		status := deleteEntry(ctx, r.Header.Get(http2.ContentLocation))
		if !status.OK() {
			e.Handle(status, runtime.RequestId(ctx), postLoc)
		}
		return nil, status
	default:
	}
	return nil, runtime.NewStatus(http.StatusMethodNotAllowed)
}

func putEntry(ctx any, variant string, body any) *runtime.Status {
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

func deleteEntry(ctx any, variant string) *runtime.Status {
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

/*
	case http.MethodGet:
		switch variant {
		case EntryV2Variant:
			entries, status := getEntry[[]EntryV2](ctx, variant, r.URL)
			if !status.OK() {
				e.Handle(status, runtime.RequestId(r), doLoc)
				return nil, status
			}
			//if len(entries) == 0 {
			//	return nil, runtime.NewStatus(http.StatusNotFound)
			//}
			return entries, runtime.NewStatusOK()
		case EntryV1Variant:
			entries, status := getEntry[[]EntryV1](ctx, variant, r.URL)
			if !status.OK() {
				e.Handle(status, runtime.RequestId(r), doLoc)
				return nil, status
			}
			//entries := queryEntriesV1(r.URL)
			//if len(entries) == 0 {
			//	return nil, runtime.NewStatus(http.StatusNotFound)
			//}
			return entries, runtime.NewStatusOK()
		default:
		}
		return nil, runtime.NewStatus(runtime.StatusInvalidContent)

*/
