package google

import (
	"fmt"
	"github.com/go-ai-agent/core/httpx"
	io2 "github.com/go-ai-agent/core/io"
	"github.com/go-ai-agent/core/log"
	"github.com/go-ai-agent/core/runtime"
	"net/http"
	"reflect"
)

type pkg struct{}

//https://www.google.com/search?q=test&rlz=1C1CHBF

var (
	Pattern = pkgPath + "/"

	PkgUri  = reflect.TypeOf(any(pkg{})).PkgPath()
	pkgPath = runtime.PathFromUri(PkgUri)

	controller         = log.NewController2(newDoHandler[runtime.LogError]())
	searchLocation     = PkgUri + "/searchHandler"
	googleQueryArgName = "q"

	// As a rule do not create/use embedded URI's, use endpoints with a sidecar like Envoy for endpoint -> URI resolution.
	// URI resolution is late and dynamic
	// With Envoy
	//googleEndpoint = "/google/search"

	// Without Envoy, this URL will pass through httpx.Resolve()
	googleEndpoint = "https://www.google.com/search"
)

// newDoHandler - templated function providing a DoHandler
func newDoHandler[E runtime.ErrorHandler]() runtime.DoHandler {
	return func(ctx any, r *http.Request, body any) (any, *runtime.Status) {
		return doHandler[E](ctx, r, body)
	}
}

func Do(ctx any, r *http.Request, body any) (any, *runtime.Status) {
	return controller.Apply(ctx, r, body)
}

func doHandler[E runtime.ErrorHandler](ctx any, r *http.Request, body any) (any, *runtime.Status) {
	if r == nil {
		return nil, runtime.NewStatus(http.StatusBadRequest)
	}
	requestId := runtime.GetOrCreateRequestId(r)
	if r.Header.Get(runtime.XRequestId) == "" {
		r.Header.Set(runtime.XRequestId, requestId)
	}
	// Need to create as new request as upstream calls may not be http, and rely on the context for a request id
	rc := r.Clone(runtime.NewRequestIdContext(r.Context(), requestId))
	switch rc.Method {
	case http.MethodGet:
		var e E

		req, err := http.NewRequest(http.MethodGet, httpx.Resolve(searchUri(rc.URL, googleEndpoint)), nil)
		if err != nil {
			return nil, e.Handle(runtime.NewStatusError(http.StatusInternalServerError, searchLocation, err), requestId, "")
		}
		// exchange.Do() will always return a non nil *http.Response
		resp, status := httpx.Do(req)
		if !status.OK() {
			return nil, e.Handle(status, requestId, searchLocation)
		}
		var buf []byte
		buf, status = io2.ReadAll(resp.Body)
		if !status.OK() {
			return nil, e.Handle(status, requestId, searchLocation)
		}
		status.Header().Set(httpx.ContentType, resp.Header.Get(httpx.ContentType))
		status.Header().Set(httpx.ContentLength, fmt.Sprintf("%v", len(buf)))
		return buf, status
	}
	return nil, runtime.NewStatus(http.StatusMethodNotAllowed)
}

func HttpHandler(w http.ResponseWriter, r *http.Request) {
	httpHandler[runtime.LogError](w, r)
}

func httpHandler[E runtime.ErrorHandler](w http.ResponseWriter, r *http.Request) *runtime.Status {
	result, status := Do(nil, r, nil)
	httpx.WriteResponse[E](w, result, status, status.Header())
	return status
}

/*
	switch r.Method {
	case http.MethodGet:
		var e E

		req, err := http.NewRequest(http.MethodGet, httpx.Resolve(searchUri(r.URL, googleEndpoint)), nil)
		if err != nil {
			status := e.Handle("", searchLocation, err).SetCode(runtime.StatusInternal)
			httpx.WriteMinResponse[E](w, status, nil)
			return status
		}
		// exchange.Do() will always return a non nil *http.Response
		resp, status := exchange.Do(req)
		if !status.OK() {
			e.HandleStatus(status, "", searchLocation)
			httpx.WriteMinResponse[E](w, status, nil)
			return status
		}
		var buf []byte
		buf, status = httpx.ReadAll(resp.Body)
		if !status.OK() {
			e.HandleStatus(status, "", searchLocation)
			httpx.WriteMinResponse[E](w, status, nil)
			return status
		}
		httpx.WriteResponse[E](w, buf, status, []httpx.Attr{{httpx.ContentType, resp.Header.Get(httpx.ContentType)}})
		return runtime.NewStatusOK()
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	return runtime.NewHttpStatus(http.StatusMethodNotAllowed)
}

*/
