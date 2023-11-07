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
	PkgUri     = reflect.TypeOf(any(pkg{})).PkgPath()
	pkgPath    = runtime.PathFromUri(PkgUri)
	controller = log.NewController(newTypeHandler[runtime.LogError]())
	typeLoc    = pkgPath + "/typeHandler"
)

// newTypeHandler - templated function providing a TypeHandlerFn with a closure
func newTypeHandler[E runtime.ErrorHandler]() runtime.TypeHandlerFn {
	return func(r *http.Request, body any) (any, *runtime.Status) {
		return typeHandler[E](r, body)
	}
}

// BodyConstraints - defining constraints for the TypeHandler body
type BodyConstraints interface {
	[]EntryV1 | []byte | runtime.Nillable
}

func TypeHandler[T BodyConstraints](r *http.Request, body T) (any, *runtime.Status) {
	return controller.Apply(httpx.UpdateHeadersAndContext(r), body)
}

func typeHandler[E runtime.ErrorHandler](r *http.Request, body any) (any, *runtime.Status) {
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
				e.Handle(status, runtime.RequestId(r), typeLoc)
				return nil, status.AddLocation(typeLoc)
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
