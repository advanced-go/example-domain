package activity

import (
	"github.com/go-ai-agent/core/http2"
	"github.com/go-ai-agent/core/io2"
	"github.com/go-ai-agent/core/json2"
	"github.com/go-ai-agent/core/log2"
	"github.com/go-ai-agent/core/runtime"
	"net/http"
	"reflect"
)

type pkg struct{}

var (
	Pattern = pkgPath + "/"

	PkgUri  = reflect.TypeOf(any(pkg{})).PkgPath()
	pkgPath = runtime.PathFromUri(PkgUri)
	loc     = pkgPath + "/entryHandler"

	wrapper        = log2.WrapDo(newDoHandler[runtime.LogError]())
	doLoc          = pkgPath + "/doHandler"
	EntryV1Variant = PkgUri + "/" + reflect.TypeOf(EntryV1{}).Name()
)

// GetConstraints - Get constraints
type GetConstraints interface {
	[]EntryV1
}

// Get - generic get function with context and uri for resource selection and filtering
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
	req, status := http2.NewRequest(ctx, method, uri, variant)
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
		// Set content-type
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
			status := json2.Unmarshal(ptr, &entries)
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

// HttpHandler - Http handler endpoint
func HttpHandler(w http.ResponseWriter, r *http.Request) {
	httpHandler[runtime.LogError](w, r)
}

func httpHandler[E runtime.ErrorHandler](w http.ResponseWriter, r *http.Request) *runtime.Status {
	if r == nil {
		w.WriteHeader(http.StatusBadRequest)
		return runtime.NewStatus(http.StatusBadRequest)
	}
	http2.AddRequestId(r)
	switch r.Method {
	case http.MethodGet:
		var buf []byte

		entries, status := Do[runtime.Nillable](r, r.Method, r.URL.String(), r.Header.Get(http2.ContentLocation), nil)
		if !status.OK() {
			http2.WriteResponse[E](w, nil, status, nil)
			return status
		}
		buf, status = json2.Marshal(entries)
		if !status.OK() {
			var e E
			e.Handle(status, runtime.RequestId(r), loc)
			http2.WriteResponse[E](w, nil, status, nil)
			return status
		}
		http2.WriteResponse[E](w, buf, status, []http2.Attr{{http2.ContentType, http2.ContentTypeJson}})
		return status
	case http.MethodPut:
		var e E

		buf, status := io2.ReadAll(r.Body)
		if !status.OK() {
			e.Handle(status, runtime.RequestId(r), loc)
			http2.WriteResponse[E](w, nil, status, nil)
			return status
		}
		_, status = Do[[]byte](r, r.Method, r.URL.String(), r.Header.Get(http2.ContentLocation), buf)
		http2.WriteResponse[E](w, nil, status, nil)
		return status
	case http.MethodDelete:
		_, status := Do[runtime.Nillable](r, r.Method, r.URL.String(), "", nil)
		http2.WriteResponse[E](w, nil, status, nil)
		return status
	default:
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	return runtime.NewStatus(http.StatusMethodNotAllowed)
}

// newDoHandler - templated function providing a DoHandler
func newDoHandler[E runtime.ErrorHandler]() runtime.DoHandler {
	return func(ctx any, r *http.Request, body any) (any, *runtime.Status) {
		return doHandler[E](ctx, r, body)
	}
}

/*
	if buf == nil {
		nc := runtime.NewStatus(runtime.StatusInvalidContent)
		http2.WriteResponse[E](w, nil, nc, nil)
		return nc
	}
	status = json2.Unmarshal(buf, &entries)
	if !status.OK() {
		e.Handle(status, requestId, loc)
	} else {
		addEntry(entries)
	}

*/

//requestId := runtime.GetOrCreateRequestId(r)
//if r.Header.Get(runtime.XRequestId) == "" {
//	r.Header.Set(runtime.XRequestId, requestId)
//}
// Handled in Http
// Need to create as new request as upstream calls may not be http, and rely on the context for a request id
//rc := r.Clone(runtime.NewRequestIdContext(r.Context(), requestId))

/*

//controller  = log2.NewController(newTypeHandler[runtime.LogError]())
// newTypeHandler - templated function providing a TypeHandlerFn
func newTypeHandler[E runtime.ErrorHandler]() runtime.TypeHandlerFn {
	return func(r *http.Request, body any) (any, *runtime.Status) {
		return doHandler[E](nil, r, body)
	}
}

*/
