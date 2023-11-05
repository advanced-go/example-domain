package timeseries

import (
	"github.com/go-ai-agent/core/resiliency"
	"github.com/go-ai-agent/core/runtime"
	"net/http"
	"reflect"
)

type pkg struct{}

var (
	PkgUri     = reflect.TypeOf(any(pkg{})).PkgPath()
	pkgPath    = runtime.PathFromUri(PkgUri)
	controller = resiliency.NewBypassController(newTypeHandler[runtime.LogError]())
)

// newTypeHandler - templated function providing a TypeHandlerFn with a closure
func newTypeHandler[E runtime.ErrorHandler]() runtime.TypeHandlerFn {
	return func(r *http.Request, body any) (any, *runtime.Status) {
		return typeHandler[E](r, body)
	}
}

// InConstraints - defining constraints for the TypeHandler
type InConstraints interface {
	[]EntryV1 | runtime.Nil
}

func TypeHandler[T InConstraints](r *http.Request, body T) (any, *runtime.Status) {
	return controller.Apply(r, body)
}

func typeHandler[E runtime.ErrorHandler](r *http.Request, body any) (any, *runtime.Status) {
	if r == nil {
		return nil, runtime.NewStatus(http.StatusBadRequest)
	}
	requestId := runtime.GetOrCreateRequestId(r)
	if r.Header.Get(runtime.XRequestId) == "" {
		r.Header.Set(runtime.XRequestId, requestId)
	}
	// Need to create as new request as upstream calls may not be http, and rely on the context for a request id
	rc := r.Clone(runtime.ContextWithRequestId(r.Context(), requestId))
	switch rc.Method {
	case http.MethodGet:
		entries := queryEntries(rc.URL)
		if len(entries) == 0 {
			return nil, runtime.NewStatus(http.StatusNotFound)
		}
		return entries, runtime.NewStatusOK()
	case http.MethodPut:
		var entries []EntryV1

		switch ptr := any(body).(type) {
		case []EntryV1:
			entries = ptr
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
