package slo

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
	Pattern = pkgPath + "/"

	PkgUri  = reflect.TypeOf(any(pkg{})).PkgPath()
	pkgPath = runtime.PathFromUri(PkgUri)
	loc     = pkgPath + "/httpHandler"

	controller = log.NewController2(newDoHandler[runtime.LogError]())
	doLoc      = pkgPath + "/doHandler"

	EntryV1Variant = PkgUri + "/" + reflect.TypeOf(EntryV1{}).Name()
)

// GetConstraints - defining constraints for Get
type GetConstraints interface {
	[]EntryV1
}

// Get - type templated function
func Get[T GetConstraints](ctx any, uri string) (T, *runtime.Status) {
	data, status := Do[runtime.Nillable](ctx, "", uri, EntryV1Variant, nil)
	if !status.OK() {
		return nil, status
	}
	if entry, ok := data.([]EntryV1); ok {
		return entry, status
	}
	return nil, runtime.NewStatus(runtime.StatusInvalidContent)
}

// BodyConstraints - defining constraints for Do
type BodyConstraints interface {
	[]EntryV1 | []byte | runtime.Nillable
}

func Do[T BodyConstraints](ctx any, method, uri, variant string, body T) (any, *runtime.Status) {
	req, status := httpx.NewRequest(ctx, method, uri, variant)
	if !status.OK() {
		return nil, status
	}
	return controller.Apply(ctx, req, body)
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

		switch ptr := any(body).(type) {
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

// HttpHandler - http handler endpoint
func HttpHandler(w http.ResponseWriter, r *http.Request) {
	httpHandler[runtime.LogError](w, r)
}

func httpHandler[E runtime.ErrorHandler](w http.ResponseWriter, r *http.Request) *runtime.Status {
	if r == nil {
		w.WriteHeader(http.StatusBadRequest)
		return runtime.NewStatus(http.StatusBadRequest)
	}
	r = httpx.UpdateHeadersAndContext(r)
	switch r.Method {
	case http.MethodGet:
		var buf []byte

		entries, status := Do[runtime.Nillable](r, r.Method, r.URL.String(), r.Header.Get(runtime.ContentLocation), nil)
		if !status.OK() {
			httpx.WriteResponse[E](w, nil, status, nil)
			return status
		}
		buf, status = json.Marshal(entries)
		if !status.OK() {
			var e E
			e.Handle(status, runtime.RequestId(r), loc)
			httpx.WriteResponse[E](w, nil, status, nil)
			return status
		}
		httpx.WriteResponse[E](w, buf, status, []httpx.Attr{{httpx.ContentType, httpx.ContentTypeJson}})
		return status
	case http.MethodPut:
		var e E

		buf, status := httpx.ReadAll(r.Body)
		if !status.OK() {
			e.Handle(status, runtime.RequestId(r), loc)
			httpx.WriteResponse[E](w, nil, status, nil)
			return status
		}
		_, status = Do[[]byte](r, r.Method, r.URL.String(), r.Header.Get(runtime.ContentLocation), buf)
		httpx.WriteResponse[E](w, nil, status, nil)
		return status
	case http.MethodDelete:
		_, status := Do[runtime.Nillable](r, r.Method, r.URL.String(), "", nil)
		httpx.WriteResponse[E](w, nil, status, nil)
		return status
	default:
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	return runtime.NewStatus(http.StatusMethodNotAllowed)
}

// newDoHandler - templated function providing a DoeHandler
func newDoHandler[E runtime.ErrorHandler]() runtime.DoHandler {
	return func(ctx any, r *http.Request, body any) (any, *runtime.Status) {
		return doHandler[E](ctx, r, body)
	}
}

/*
	if buf == nil {
		nc := runtime.NewStatus(runtime.StatusInvalidContent)
		httpx.WriteResponse[E](w, nil, nc, nil)
		return nc
	}
	status = json.Unmarshal(buf, &entries)
	if !status.OK() {
		e.Handle(status, requestId, loc)
	} else {
		addEntry(entries)
	}
*/
