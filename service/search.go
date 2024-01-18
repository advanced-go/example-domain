package service

import (
	"github.com/advanced-go/core/exchange"
	"github.com/advanced-go/core/http2"
	"github.com/advanced-go/core/runtime"
	"net/http"
	"strings"
)

const (
// searchTemplate = "github.com/advanced-go/search/provider:search?%v"
// http://localhost:8081/github.com/advanced-go/search/provider:search?q=golang
)

func searchHandler[E runtime.ErrorHandler](w http.ResponseWriter, r *http.Request) runtime.Status {
	if r == nil {
		w.WriteHeader(http.StatusBadRequest)
		return runtime.NewStatus(http.StatusBadRequest)
	}
	switch strings.ToUpper(r.Method) {
	case http.MethodGet:
		newUrl := resolver.Build(searchTemplate, r.URL.Query().Encode())
		resp, status := exchange.Get(newUrl, r.Header)
		if !status.OK() {
			http2.WriteResponse[E](w, nil, status, nil)
		} else {
			http2.WriteResponse[E](w, resp.Body, status, resp.Header)
		}
		return status
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return runtime.NewStatus(http.StatusMethodNotAllowed)
	}
}
