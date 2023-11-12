package timeseries

import (
	"github.com/go-ai-agent/core/http2"
	"github.com/go-ai-agent/core/runtime"
	"net/http"
	"reflect"
)

var (
	Pattern = pkgPath + "/"

	PkgUri  = reflect.TypeOf(any(pkg{})).PkgPath()
	pkgPath = runtime.PathFromUri(PkgUri)

	EntryV1Variant = PkgUri + "/" + reflect.TypeOf(EntryV1{}).Name()
	EntryV2Variant = PkgUri + "/" + reflect.TypeOf(EntryV2{}).Name()
)

// GetConstraints - Get constraints
type GetConstraints interface {
	[]EntryV1 | []EntryV2
}

// Get - generic get function with context and uri for resource selection and filtering
func Get[T GetConstraints](ctx any, uri string) (T, *runtime.Status) {
	var t T

	switch ptr := any(&t).(type) {
	case *[]EntryV1:
		data, status := Do(ctx, "", uri, EntryV1Variant, nil)
		if !status.OK() {
			return nil, status
		}
		if entry, ok := data.([]EntryV1); ok {
			*ptr = entry
		}
	case *[]EntryV2:
		data, status := Do(ctx, "", uri, EntryV2Variant, nil)
		if !status.OK() {
			return nil, status
		}
		if entry, ok := data.([]EntryV2); ok {
			*ptr = entry
		}
	default:
		return nil, runtime.NewStatus(runtime.StatusInvalidContent)
	}
	return t, runtime.NewStatusOK()
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
	return doWrapper(ctx, req, body)
}

// HttpHandler - http endpoint
func HttpHandler(w http.ResponseWriter, r *http.Request) {
	httpWrapper(nil, w, r) //httpHandler[runtime.LogError](nil, w, r)
}
