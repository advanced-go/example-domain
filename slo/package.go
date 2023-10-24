package slo

import (
	"encoding/json"
	"github.com/go-ai-agent/core/httpx"
	"github.com/go-ai-agent/core/runtime"
	"net/http"
	"reflect"
	"sync/atomic"
)

type pkg struct{}

var (
	PkgUrl  = runtime.ParsePkgUrl(reflect.TypeOf(any(pkg{})).PkgPath())
	PkgUri  = PkgUrl.Host + PkgUrl.Path
	SLOPath = PkgUrl.Path + "/entry"
	started int64
)

// IsPkgStarted - returns status of startup
func IsPkgStarted() bool {
	return atomic.LoadInt64(&started) != 0
}

func DoHandler(req *http.Request) (*http.Response, error) {
	recorder := httpx.NewRecorder()
	status := sloHandler[runtime.BypassError](recorder, req)
	var err error
	if status.IsErrors() {
		err = status.Errors()[0]
	}
	return recorder.Result(), err
}

func SLOHandler(w http.ResponseWriter, r *http.Request) {
	sloHandler[runtime.LogError](w, r)
}

func sloHandler[E runtime.ErrorHandler](w http.ResponseWriter, r *http.Request) *runtime.Status {
	if r == nil {
		w.WriteHeader(http.StatusBadRequest)
		return runtime.NewHttpStatusCode(http.StatusBadRequest)
	}
	switch r.Method {
	case "GET":
		buf, status := marshalSLO[E](GetSLO())
		httpx.WriteResponse[E](w, buf, status)
		return status
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
	return runtime.NewStatusOK()
}

func marshalSLO[E runtime.ErrorHandler](entry []SLO) ([]byte, *runtime.Status) {
	buf, err := json.Marshal(entry)
	if err != nil {
		var e E
		return nil, e.Handle(nil, "marshal", err)
	}
	return buf, runtime.NewStatusOK()
}

func unmarshalSLO[E runtime.ErrorHandler](buf []byte) ([]SLO, *runtime.Status) {
	var entry []SLO

	err := json.Unmarshal(buf, &entry)
	if err != nil {
		var e E
		return nil, e.Handle(nil, "unmarshal", err)
	}
	return entry, runtime.NewStatusOK()
}
