package activity

import (
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
	requestId := runtime.GetOrCreateRequestId(r)
	if r.Header.Get(runtime.XRequestId) == "" {
		r.Header.Set(runtime.XRequestId, requestId)
	}
	// Need to create as new request as upstream calls may not be http, and rely on the context for a request id
	rc := r.Clone(runtime.ContextWithRequestId(r.Context(), requestId))
	switch rc.Method {
	case http.MethodGet:
		e := queryEntries(rc)
		if len(e) == 0 {
			status := runtime.NewStatusCode(runtime.StatusNotFound)
			httpx.WriteMinResponse[E](w, status)
			return status
		}
		buf, status := runtime.MarshalType(e)
		status.SetRequestId(requestId)
		if !status.OK() {
			httpx.WriteMinResponse[E](w, status)
			return status
		}
		httpx.WriteResponse[E](w, buf, status, runtime.ContentType, runtime.ContentTypeJson)
		return status
	case http.MethodPut:
		var entries []entry

		buf, status := httpx.ReadAll(rc.Body)
		if !status.OK() || buf == nil {
			httpx.WriteMinResponse[E](w, status.SetRequestId(requestId))
			return status
		}
		entries, status = runtime.UnmarshalType[[]entry](buf)
		if status.OK() {
			addEntry(entries)
		}
		httpx.WriteMinResponse[E](w, status.SetRequestId(requestId))
		return status
	case http.MethodDelete:
		deleteEntries()
		status := runtime.NewStatusOK()
		httpx.WriteMinResponse[E](w, status.SetRequestId(requestId))
		return status
	default:
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	return runtime.NewHttpStatusCode(http.StatusMethodNotAllowed)
}

func queryEntries(r *http.Request) []entry {
	var result []entry

	name := ""
	if r.URL.Query() != nil {
		name = r.URL.Query().Get(Type)
	}
	if len(name) != 0 {
		result = getEntriesByType(name)
	} else {
		result = getEntries()
	}
	return result
}

/*
func entryHandler[E runtime.ErrorHandler](w http.ResponseWriter, r *http.Request) *runtime.Status {
	if r == nil {
		w.WriteHeader(http.StatusBadRequest)
		return runtime.NewHttpStatusCode(http.StatusBadRequest)
	}

	return runtime.NewStatusOK()
}


func MarshalEntry[E runtime.ErrorHandler](entry []Entry) ([]byte, *runtime.Status) {
	buf, err := json.Marshal(entry)
	if err != nil {
		var e E
		return nil, e.Handle(nil, "marshal", err)
	}
	return buf, runtime.NewStatusOK()
}

func UnmarshalEntry[E runtime.ErrorHandler](buf []byte) ([]Entry, *runtime.Status) {
	var entry []Entry

	err := json.Unmarshal(buf, &entry)
	if err != nil {
		var e E
		return nil, e.Handle(nil, "unmarshal", err)
	}
	return entry, runtime.NewStatusOK()
}


*/
