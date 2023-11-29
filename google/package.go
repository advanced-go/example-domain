package google

import (
	"github.com/advanced-go/core/runtime"
	"net/http"
)

type pkg struct{}

// https://www.google.com/search?q=test&rlz=1C1CHBF

const (
	PkgPath = "github.com/advanced-go/example-domain/google"
	Pattern = "/" + PkgPath + "/"
)

const (
	searchLocation     = PkgPath + ":searchHandler"
	googleQueryArgName = "q"

	// As a rule do not create/use embedded URI's, use endpoints with a sidecar like Envoy for endpoint -> URI resolution.
	// URI resolution is late and dynamic
	// With Envoy
	//googleEndpoint = "/google/search"

	// Without Envoy, this URL will pass through httpx.Resolve()
	googleEndpoint = "https://www.google.com/search"
)

// Get - exchange handler
func Get(r *http.Request) (any, runtime.Status) {
	return getHandler[runtime.LogError](r)
}

// HttpHandler - HTTP handler endpoint
func HttpHandler(w http.ResponseWriter, r *http.Request) {
	httpHandler[runtime.LogError](w, r)
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
		return runtime.StatusOK()
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	return runtime.NewHttpStatus(http.StatusMethodNotAllowed)
}

*/
