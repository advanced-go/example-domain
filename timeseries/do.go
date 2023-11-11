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
	doLoc   = pkgPath + "/doHandler"
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
	variant := r.Header.Get(http2.ContentLocation)
	if variant != EntryV1Variant && variant != EntryV2Variant {
		return nil, runtime.NewStatus(http.StatusBadRequest)
	}
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
		var entries []EntryV1

		switch ptr := body.(type) {
		case []EntryV1:
			entries = ptr
		case []byte:
			if ptr == nil {
				return nil, runtime.NewStatus(runtime.StatusInvalidContent)
			}
			status := json2.Unmarshal(ptr, &entries)
			if !status.OK() {
				var e E
				e.Handle(status, runtime.RequestId(r), doLoc)
				return nil, status.AddLocation(doLoc)
			}
		default:
			var e E
			status := runtime.NewStatusError(runtime.StatusInvalidContent, doLoc, runtime.NewInvalidBodyTypeError(body))
			e.Handle(status, runtime.RequestId(r), "")
			return nil, status
		}
		if len(entries) == 0 {
			return nil, runtime.NewStatus(runtime.StatusInvalidContent)
		}
		addEntryV1(entries)
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
