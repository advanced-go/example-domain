package timeseries

import (
	"github.com/go-ai-agent/core/http2"
	"github.com/go-ai-agent/core/runtime"
	"net/http"
	"reflect"
)

var (
	PkgUri  = reflect.TypeOf(any(pkg{})).PkgPath()
	pkgPath = runtime.PathFromUri(PkgUri)

	EntryV1Variant = PkgUri + "/" + reflect.TypeOf(EntryV1{}).Name()
	EntryV2Variant = PkgUri + "/" + reflect.TypeOf(EntryV2{}).Name()
)

// GetConstraints - Get constraints
type GetConstraints interface {
	[]EntryV1
}

// Get - generic get function with context and uri for resource selection and filtering
func Get[T GetConstraints](ctx any, uri string) (T, *runtime.Status) {
	var t T

	switch any(t).(type) {
	case []EntryV1:
		data, status := Do(ctx, "", uri, EntryV1Variant, nil)
		if !status.OK() {
			return nil, status
		}
		if entry, ok := data.([]EntryV1); ok {
			return entry, status
		}
	default:
	}
	return nil, runtime.NewStatus(runtime.StatusInvalidContent)
}

// DoConstraints - Do constraints
type DoConstraints interface {
	[]EntryV1 | []byte | runtime.Nillable
}

// Do - exchange function
func Do(ctx any, method, uri, variant string, body any) (any, *runtime.Status) {
	req, status := http2.NewRequest(ctx, method, uri, variant)
	if !status.OK() {
		return nil, status
	}
	return wrapper(ctx, req, body)
}

// HttpHandler - http handler endpoint
func HttpHandler(w http.ResponseWriter, r *http.Request) {
	httpHandler[runtime.LogError](nil, w, r)
}
