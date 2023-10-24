package timeseries

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
	PkgUrl    = runtime.ParsePkgUrl(reflect.TypeOf(any(pkg{})).PkgPath())
	PkgUri    = PkgUrl.Host + PkgUrl.Path
	EntryPath = PkgUrl.Path + "/entry"
	started   int64
)

// IsPkgStarted - returns status of startup
func IsPkgStarted() bool {
	return atomic.LoadInt64(&started) != 0
}

func DoHandler(req *http.Request) (*http.Response, error) {
	recorder := httpx.NewRecorder()
	status := entryHandler[runtime.BypassError](recorder, req)
	var err error
	if status.IsErrors() {
		err = status.Errors()[0]
	}
	return recorder.Result(), err
}

func EntryHandler(w http.ResponseWriter, r *http.Request) {
	entryHandler[runtime.LogError](w, r)
}

func entryHandler[E runtime.ErrorHandler](w http.ResponseWriter, r *http.Request) *runtime.Status {
	if r == nil {
		w.WriteHeader(http.StatusBadRequest)
		return runtime.NewHttpStatusCode(http.StatusBadRequest)
	}
	switch r.Method {
	case http.MethodGet:
		var entries []entry

		name := ""
		if r.URL.Query() != nil {
			name = r.URL.Query().Get(ConrollerName)
		}
		if len(name) != 0 {
			entries = getEntriesByController(name)
		} else {
			entries = getEntries()
		}
		buf, status := marshalEntry[E](entries)
		//if status.OK() {
		//	status.SetMetadata(runtime.ContentType, runtime.ContentTypeJson)
		//} else {
		//	status.SetMetadata(runtime.ContentType, runtime.ContentTypeText)
		//}
		httpx.WriteResponse[E](w, buf, status)
		return status
	case http.MethodPut:
		buf, status := httpx.ReadAll[E](r.Body)
		if !status.OK() {
			httpx.WriteResponse[E](w, nil, status)
			return status
		}
		entries, status1 := unmarshalEntry[E](buf)
		if status.OK() {
			addEntry(entries)
		}
		httpx.WriteResponse[E](w, nil, status1)
		return status1
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
	return runtime.NewStatusOK()
}

func marshalEntry[E runtime.ErrorHandler](entry []entry) ([]byte, *runtime.Status) {
	buf, err := json.Marshal(entry)
	if err != nil {
		var e E
		return nil, e.Handle(nil, "marshal", err)
	}
	return buf, runtime.NewStatusOK()
}

func unmarshalEntry[E runtime.ErrorHandler](buf []byte) ([]entry, *runtime.Status) {
	var e []entry

	err := json.Unmarshal(buf, &e)
	if err != nil {
		var e E
		return nil, e.Handle(nil, "unmarshal", err)

	}
	return e, runtime.NewStatusOK()
}
