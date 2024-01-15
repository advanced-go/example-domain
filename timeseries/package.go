package timeseries

import (
	"github.com/advanced-go/core/runtime"
	"github.com/advanced-go/example-domain/timeseries/entryv1"
	"github.com/advanced-go/example-domain/timeseries/entryv2"
	"net/http"
	"net/url"
)

const (
	PkgPath         = "github/advanced-go/example-domain/timeseries"
	v1EntryResource = "v1/entry"
	v2EntryResource = "v2/entry"
)

// GetEntryV1 - get entries
func GetEntryV1(h http.Header, values url.Values) (entries []entryv1.Entry, status runtime.Status) {
	return entryv1.Get(h, values)
}

// GetEntryV2 - get entries
func GetEntryV2(h http.Header, values url.Values) (entries []entryv2.Entry, status runtime.Status) {
	return entryv2.Get(h, values)
}

// PostEntryV1 - exchange function
func PostEntryV1[T entryv1.PostConstraints](h http.Header, method string, values url.Values, body T) (t any, status runtime.Status) {
	return entryv1.Post[T](h, method, values, body)
}

// PostEntryV2 - exchange function
func PostEntryV2[T entryv2.PostConstraints](h http.Header, method string, values url.Values, body T) (t any, status runtime.Status) {
	return entryv2.Post[T](h, method, values, body)
}

/*
// HttpHandler - http endpoint
func HttpHandler(w http.ResponseWriter, r *http.Request) {
	_, rsc, ok := uri.UprootUrn(r.URL.Path)
	if !ok {
		status := runtime.NewStatus(http.StatusBadRequest)
		status.SetContent(errors.New(fmt.Sprintf("error invalid path, not a valid URN: %v", r.URL.Path)), false)
		http2.WriteResponse[runtime.Log](w, nil, status, nil)
		return
	}
	runtime.AddRequestId(r)
	switch strings.ToLower(rsc) {
	case v1EntryResource:
		entryv1.HttpHandler(w, r)
	case v2EntryResource:
		entryv2.HttpHandler(w, r)
	default:
		status := runtime.NewStatus(http.StatusNotFound)
		status.SetContent(errors.New(fmt.Sprintf("error invalid URI, resource was not found: %v", rsc)), false)
		http2.WriteResponse[runtime.Log](w, nil, status, nil)
	}
}


*/
