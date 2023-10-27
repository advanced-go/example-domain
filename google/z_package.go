package google

import (
	"errors"
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

	pkgUri  = reflect.TypeOf(any(pkg{})).PkgPath()
	pkgPath = runtime.PathFromUri(pkgUri)

	searchLocation = pkgUri + "/searchHandler"
	queryArgName   = "q"
	searchPath     = "/search"
)

// IsPkgStarted - returns status of startup
func IsPkgStarted() bool {
	return true
}

func DoHandler(req *http.Request) (*http.Response, error) {
	recorder := httpx.NewRecorder()
	status := searchHandler[runtime.BypassError](recorder, req)
	var err error
	if status.IsErrors() {
		err = status.FirstError()
	}
	return recorder.Result(), err
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

		uri := exchange.Resolve(searchEndpoint(r.URL))
		req, err := http.NewRequest(http.MethodGet, uri, nil)
		if err != nil {
			status := runtime.NewStatusError(runtime.StatusInternal, searchLocation, err)
			e.HandleStatus(status, "")
			httpx.WriteMinResponse[E](w, status)
			return status
		}
		resp, status := exchange.Do[E](req)
		if !status.OK() {
			e.HandleStatus(status, "")
			httpx.WriteMinResponse[E](w, status)
			return status
		}
		if resp == nil {
			rn := runtime.NewStatusError(runtime.StatusInternal, searchLocation, errors.New("error: response is nil"))
			e.HandleStatus(rn, "")
			httpx.WriteMinResponse[E](w, rn)
			return rn
		}
		var buf []byte
		buf, status = httpx.ReadAll(resp.Body)
		if !status.OK() {
			e.HandleStatus(status, "")
			httpx.WriteMinResponse[E](w, status)
			return status
		}
		httpx.WriteResponse[E](w, buf, status, httpx.ContentType, resp.Header.Get(httpx.ContentType))
		return runtime.NewStatusOK()
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	return runtime.NewHttpStatus(http.StatusMethodNotAllowed)
}
