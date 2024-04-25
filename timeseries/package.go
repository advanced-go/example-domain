package timeseries

import (
	"github.com/advanced-go/example-domain/timeseries/entryv1"
	"github.com/advanced-go/example-domain/timeseries/entryv2"
	"github.com/advanced-go/stdlib/core"
	"net/http"
	"net/url"
)

const (
	PkgPath = "github/advanced-go/example-domain/timeseries"
)

// GetEntryV1 - get entries
func GetEntryV1(h http.Header, values url.Values) (entries []entryv1.Entry, status *core.Status) {
	return entryv1.Get(h, values)
}

// GetEntryV2 - get entries
func GetEntryV2(h http.Header, values url.Values) (entries []entryv2.Entry, status *core.Status) {
	return entryv2.Get(h, values)
}

// PostEntryV1 - exchange function
func PostEntryV1[T entryv1.PostConstraints](h http.Header, method string, values url.Values, body T) (t any, status *core.Status) {
	return entryv1.Post[T](h, method, values, body)
}

// PostEntryV2 - exchange function
func PostEntryV2[T entryv2.PostConstraints](h http.Header, method string, values url.Values, body T) (t any, status *core.Status) {
	return entryv2.Post[T](h, method, values, body)
}
