package service

import (
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"net/http"
	"strings"
)

const (
// searchTemplate = "github.com/advanced-go/search/provider:search?%v"
// http://localhost:8081/github.com/advanced-go/search/provider:search?q=golang
)

func searchHandler[E core.ErrorHandler](w http.ResponseWriter, r *http.Request) *core.Status {
	if r == nil {
		w.WriteHeader(http.StatusBadRequest)
		return core.NewStatus(http.StatusBadRequest)
	}
	switch strings.ToUpper(r.Method) {
	case http.MethodGet:
		newUrl := resolver.Build(searchTemplate, r.URL.Query().Encode())
		resp, status := httpx.Get(nil, newUrl, r.Header)
		if !status.OK() {
			httpx.WriteResponse[E](w, nil, status.HttpCode(), nil)
		} else {
			httpx.WriteResponse[E](w, resp.Header, status.HttpCode(), resp.Body)
		}
		return status
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return core.NewStatus(http.StatusMethodNotAllowed)
	}
}
