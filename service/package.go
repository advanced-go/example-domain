package service

import (
	"errors"
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"net/http"
	"strings"
)

//type pkg struct{}

const (
	PkgPath          = "github/advanced-go/example-domain/service"
	activityPath     = "activity/entry"
	sloPath          = "slo/entry"
	timeseriesPathV1 = "timeseries/v1/entry"
	timeseriesPathV2 = "timeseries/v2/entry"
	searchPath       = "search"
)

// HttpHandler - Http endpoint
func HttpHandler(w http.ResponseWriter, r *http.Request) {
	path, status0 := httpx.ValidateRequest(r, PkgPath)
	if !status0.OK() {
		httpx.WriteResponse[core.Log](w, nil, status0.HttpCode(), status0.Err)
		return
	}
	core.AddRequestId(r)
	switch strings.ToLower(path) {
	case activityPath:
		activityHandler[core.Log](w, r)
	case sloPath:
		sloHandler[core.Log](w, r)
	case timeseriesPathV1:
		timeseriesHandlerV1[core.Log](w, r)
	case timeseriesPathV2:
		timeseriesHandlerV2[core.Log](w, r)
	case searchPath:
		searchHandler[core.Log](w, r)
	default:
		status := core.NewStatusError(http.StatusNotFound, errors.New(fmt.Sprintf("error invalid URI, resource was not found: %v", path)))
		httpx.WriteResponse[core.Log](w, nil, status.HttpCode(), status.Err)
	}
}
