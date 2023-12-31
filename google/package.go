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

	searchLocation     = PkgPath + ":searchHandler"
	googleQueryArgName = "q"

	// Resolution accepts identifiers that can be tags, paths, or complete URLs.
	// With Envoy
	//googleEndpoint = "/google/search"

	// Without Envoy, this URL will pass through the resolver
	googleEndpoint = "https://www.google.com/search"
)

// Search - search handler
func Search(r *http.Request) (any, runtime.Status) {
	return searchHandler[runtime.Log](r)
}

// HttpHandler - HTTP handler endpoint
func HttpHandler(w http.ResponseWriter, r *http.Request) {
	httpHandler[runtime.Log](w, r)
}
