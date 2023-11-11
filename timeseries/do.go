package timeseries

import (
	"github.com/go-ai-agent/core/http2"
	"github.com/go-ai-agent/core/json2"
	"github.com/go-ai-agent/core/log2"
	"github.com/go-ai-agent/core/runtime"
	"net/http"
)

type pkg struct{}

var (
	wrapper = log2.WrapDo(newDoHandler[runtime.LogError]())
	doLoc   = PkgUri + "/doHandler"
	putLoc  = PkgUri + "/put"
	getLoc  = PkgUri + "/get"
)

func doHandler[E runtime.ErrorHandler](ctx any, r *http.Request, body any) (any, *runtime.Status) {
	if r == nil {
		return nil, runtime.NewStatus(http.StatusBadRequest)
	}
	if runtime.IsDebugEnvironment() {
		if fn := http2.DoHandlerProxy(ctx); fn != nil {
			return fn(ctx, r, body)
		}
	}
	variant := verifyVariant(r)
	switch r.Method {
	case http.MethodGet:
		switch variant {
		case EntryV2Variant:
			entries := queryEntriesV2(r.URL)
			if len(entries) == 0 {
				return nil, runtime.NewStatus(http.StatusNotFound)
			}
			return entries, runtime.NewStatusOK()
		case EntryV1Variant:
			entries := queryEntriesV1(r.URL)
			if len(entries) == 0 {
				return nil, runtime.NewStatus(http.StatusNotFound)
			}
			return entries, runtime.NewStatusOK()
		default:
		}
		return nil, runtime.NewStatus(runtime.StatusInvalidContent)
	case http.MethodPut:
		status := put(r.Header.Get(http2.ContentLocation), body)
		if !status.OK() {
			var e E
			e.Handle(status, runtime.RequestId(r), doLoc)
			return nil, status
		}
		return nil, runtime.NewStatusOK()
	case http.MethodDelete:
		switch variant {
		case EntryV1Variant:
			deleteEntriesV1()
		case EntryV2Variant:
			deleteEntriesV2()
		}
		return nil, runtime.NewStatusOK()
	default:
	}
	return nil, runtime.NewStatus(http.StatusMethodNotAllowed)
}

// newDoHandler - templated function providing DoHandler
func newDoHandler[E runtime.ErrorHandler]() runtime.DoHandler {
	return func(ctx any, r *http.Request, body any) (any, *runtime.Status) {
		return doHandler[E](ctx, r, body)
	}
}

func put(variant string, body any) *runtime.Status {
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
			if ptr == nil {
				return runtime.NewStatus(runtime.StatusInvalidContent)
			}
			status := json2.Unmarshal(ptr, &entries)
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
			if ptr == nil {
				return runtime.NewStatus(runtime.StatusInvalidContent)
			}
			status := json2.Unmarshal(ptr, &entries)
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

func verifyVariant(r *http.Request) string {
	variant := r.Header.Get(http2.ContentLocation)
	if len(variant) > 0 {
		return variant
	}
	v := r.URL.Query().Get("v")
	if len(v) > 0 {
		if v == "v1" {
			return EntryV1Variant
		}
		if v == "v2" {
			return EntryV2Variant
		}
	}
	return EntryV1Variant
}
