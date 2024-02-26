package service

import (
	"errors"
	"fmt"
	"github.com/advanced-go/core/http2"
	"github.com/advanced-go/core/runtime"
	"net/http"
	"strings"
)

type pkg struct{}

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
	path, status0 := http2.ValidateRequest(r, PkgPath)
	if !status0.OK() {
		http2.WriteResponse[runtime.Log](w, status0.Error(), status0, nil)
		return
	}
	runtime.AddRequestId(r)
	switch strings.ToLower(path) {
	case activityPath:
		activityHandler[runtime.Log](w, r)
	case sloPath:
		sloHandler[runtime.Log](w, r)
	case timeseriesPathV1:
		timeseriesHandlerV1[runtime.Log](w, r)
	case timeseriesPathV2:
		timeseriesHandlerV2[runtime.Log](w, r)
	case searchPath:
		searchHandler[runtime.Log](w, r)
	default:
		status := runtime.NewStatusError(http.StatusNotFound, errors.New(fmt.Sprintf("error invalid URI, resource was not found: %v", path)), nil)
		http2.WriteResponse[runtime.Log](w, status.Error(), status, nil)
	}
}
