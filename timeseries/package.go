package timeseries

import (
	"github.com/go-ai-agent/core/http2"
	"github.com/go-ai-agent/core/log2"
	"github.com/go-ai-agent/core/runtime"
	"net/http"
	"reflect"
	"time"
)

var (
	Pattern = pkgPath + "/"

	PkgUri  = reflect.TypeOf(any(pkg{})).PkgPath()
	pkgPath = runtime.PathFromUri(PkgUri)

	putLoc2 = PkgUri + "/Put"
	getLoc2 = PkgUri + "/Get"

	EntryV1Variant = PkgUri + "/" + reflect.TypeOf(EntryV1{}).Name()
	EntryV2Variant = PkgUri + "/" + reflect.TypeOf(EntryV2{}).Name()
)

func NewStatusCode(status **runtime.Status) func() int {
	return func() int {
		return (*(status)).Code()
	}
}

// GetConstraints - Get constraints
type GetConstraints interface {
	[]EntryV1 | []EntryV2 | []byte
}

func log(ctx any, method string, uri any, statusCode func() int) func() {
	start := time.Now().UTC()
	req, _ := http2.NewRequest(ctx, method, uri, "")
	return func() {
		log2.InternalAccess(start, time.Since(start), req, &http.Response{StatusCode: statusCode()}, -1, "")
	}
}

// Get - generic get function with context and uri for resource selection and filtering
func Get[T GetConstraints](ctx, uri any, variant string) (t T, status *runtime.Status) {
	defer log(ctx, "GET", uri, NewStatusCode(&status))

	t, status = getEntry[T](ctx, uri, variant)
	if !status.OK() {
		var e runtime.LogError
		e.Handle(status, runtime.RequestId(ctx), getLoc2)
	}
	return t, runtime.NewStatusOK()
}

// PostConstraints - Post constraints
type PostConstraints interface {
	[]EntryV1 | []byte | runtime.Nillable
}

// Post - exchange function
func Post(ctx any, method, uri, variant string, body any) (any, *runtime.Status) {
	req, status := http2.NewRequest(ctx, method, uri, variant)
	if !status.OK() {
		return nil, status
	}
	return postWrapper(ctx, req, body)
}

// HttpHandler - http endpoint
func HttpHandler(w http.ResponseWriter, r *http.Request) {
	httpWrapper(nil, w, r)
}
