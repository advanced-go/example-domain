package search

import (
	"errors"
	"fmt"
	"github.com/advanced-go/core/exchange"
	"github.com/advanced-go/core/http2"
	"github.com/advanced-go/core/runtime"
	"github.com/advanced-go/core/uri"
	"net/http"
)

type pkg struct{}

const (
	PkgPath = "github.com/advanced-go/example-domain/search"
	//Pattern = "/" + PkgPath + "/"

	authority  = "localhost:8081"
	searchPath = "github.com/advanced-go/search/provider:search?%v"
	//searchResource       = "search"
	//httpHandlerRouteName = "http-handler"
	postRouteName = "post-entry"
	postEntryLoc  = PkgPath + ":PostEntry"

	getRouteName = "get-entry"
	getEntryLoc  = PkgPath + ":GetEntry"
)

// HttpHandler - Http endpoint
func HttpHandler(w http.ResponseWriter, r *http.Request) {
	if r == nil {
		http2.WriteResponse[runtime.Log](w, nil, runtime.NewStatus(runtime.StatusInvalidArgument), nil)
		return
	}
	nid, rsc, ok := uri.UprootUrn(r.URL.Path)
	if !ok || nid != PkgPath {
		status := runtime.NewStatusWithContent(http.StatusBadRequest, errors.New(fmt.Sprintf("error invalid path, not a valid URN: %v", r.URL.Path)), false)
		http2.WriteResponse[runtime.Log](w, nil, status, nil)
		return
	}
	if len(rsc) > 0 {
		status := runtime.NewStatusWithContent(http.StatusNotFound, errors.New(fmt.Sprintf("error invalid URI, resource was not found: %v", rsc)), false)
		http2.WriteResponse[runtime.Log](w, nil, status, nil)
		return
	}
	runtime.AddRequestId(r)
	newUrl := resolver.Build(authority, searchPath, r.URL.Query())
	resp, status := exchange.Get(newUrl, r.Header)
	if !status.OK() {
		http2.WriteResponse[runtime.Log](w, nil, status, nil)
		return
	}
	var buf []byte
	buf, status = runtime.NewBytes(resp)
	if !status.OK() {
		http2.WriteResponse[runtime.Log](w, nil, status, nil)
		return
	}
	http2.WriteResponse[runtime.Log](w, buf, status, nil)
	//defer access.LogDeferred(access.InternalTraffic, r, httpHandlerRouteName, "", -1, "", &status)()

}
