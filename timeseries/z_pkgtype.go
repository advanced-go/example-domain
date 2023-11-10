package timeseries

import (
	"github.com/go-ai-agent/core/httpx"
	"github.com/go-ai-agent/core/json"
	"github.com/go-ai-agent/core/log"
	"github.com/go-ai-agent/core/runtime"
	"net/http"
	"reflect"
)

type pkg struct{}

var (
	PkgUri         = reflect.TypeOf(any(pkg{})).PkgPath()
	pkgPath        = runtime.PathFromUri(PkgUri)
	wrapper        = log.WrapDo(newDoHandler[runtime.LogError]())
	doLoc          = pkgPath + "/doHandler"
	EntryV1Variant = PkgUri + "/" + reflect.TypeOf(EntryV1{}).Name()
)

// GetConstraints - Get constraints
type GetConstraints interface {
	[]EntryV1
}

// Get - generic get function
func Get[T GetConstraints](ctx any, uri string) (T, *runtime.Status) {
	var t T
	//Set variant based on generic type
	variant := EntryV1Variant
	switch any(t).(type) {
	case []EntryV1:
	}
	data, status := Do[runtime.Nillable](ctx, "", uri, variant, nil)
	if !status.OK() {
		return nil, status
	}
	if entry, ok := data.([]EntryV1); ok {
		return entry, status
	}
	return nil, runtime.NewStatus(runtime.StatusInvalidContent)
}

// DoConstraints - Do constraints
type DoConstraints interface {
	[]EntryV1 | []byte | runtime.Nillable
}

// Do - generic exchange function
func Do[T DoConstraints](ctx any, method, uri, variant string, body T) (any, *runtime.Status) {
	req, status := httpx.NewRequest(ctx, method, uri, variant)
	if !status.OK() {
		return nil, status
	}
	return wrapper(ctx, req, body)
}

func doHandler[E runtime.ErrorHandler](ctx any, r *http.Request, body any) (any, *runtime.Status) {
	if r == nil {
		return nil, runtime.NewStatus(http.StatusBadRequest)
	}
	switch r.Method {
	case http.MethodGet:
		entries := queryEntries(r.URL)
		if len(entries) == 0 {
			return nil, runtime.NewStatus(http.StatusNotFound)
		}
		return entries, runtime.NewStatusOK()
	case http.MethodPut:
		var entries []EntryV1

		switch ptr := body.(type) {
		case []EntryV1:
			entries = ptr
		case []byte:
			if ptr == nil {
				return nil, runtime.NewStatus(runtime.StatusInvalidContent)
			}
			status := json.Unmarshal(ptr, &entries)
			if !status.OK() {
				var e E
				e.Handle(status, runtime.RequestId(r), doLoc)
				return nil, status.AddLocation(doLoc)
			}
		default:
		}
		if len(entries) == 0 {
			return nil, runtime.NewStatus(runtime.StatusInvalidContent)
		}
		addEntry(entries)
		return nil, runtime.NewStatusOK()
	case http.MethodDelete:
		deleteEntries()
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
