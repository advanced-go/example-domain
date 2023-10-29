package google

import (
	"github.com/go-ai-agent/core/exchange"
	"github.com/go-ai-agent/core/httpx"
	"github.com/go-ai-agent/core/runtime"
	"net/http"
	"reflect"
)

type pkg struct{}

//https://www.google.com/search?q=test&rlz=1C1CHBF

var (
	SearchEndpoint = pkgPath + "/search"

	PkgUri  = reflect.TypeOf(any(pkg{})).PkgPath()
	pkgPath = runtime.PathFromUri(PkgUri)

	searchLocation     = PkgUri + "/searchHandler"
	googleQueryArgName = "q"

	// As a rule do not create/use embedded URI's, use endpoints with a sidecar like Envoy for endpoint -> URI resolution.
	// URI resolution is late and dynamic
	// With Envoy
	//googleEndpoint = "/google/search"

	// Without Envoy, this URL will pass through httpx.Resolve()
	googleEndpoint = "https://www.google.com/search"
)

// IsPkgStarted - returns status of startup
func IsPkgStarted() bool {
	return true
}

func DoHandler(req *http.Request) (*http.Response, error) {
	w := httpx.NewRecorder()
	status := searchHandler[runtime.BypassError](w, req)
	// kludge
	w.Result().Header = w.Header().Clone()
	var err error
	if status.IsErrors() {
		err = status.FirstError()
	}
	return w.Result(), err
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	searchHandler[runtime.LogError](w, r)
}

func searchHandler[E runtime.ErrorHandler](w http.ResponseWriter, r *http.Request) *runtime.Status {
	if r == nil {
		w.WriteHeader(http.StatusBadRequest)
		return runtime.NewHttpStatus(http.StatusBadRequest)
	}
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
