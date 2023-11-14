package activity

import (
	"github.com/go-ai-agent/core/http2"
	"github.com/go-ai-agent/core/log2"
	"github.com/go-ai-agent/core/runtime"
	"net/http"
	"net/url"
	"reflect"
)

type pkg struct{}

var (
	Pattern = pkgPath + "/"

	PkgUri        = reflect.TypeOf(any(pkg{})).PkgPath()
	pkgPath       = runtime.PathFromUri(PkgUri)
	getEntryLoc2  = PkgUri + "/GetEntry"
	postEntryLoc2 = PkgUri + "/PostEntry"

	EntryV1Variant = PkgUri + "/" + reflect.TypeOf(EntryV1{}).Name()
)

// GetEntryConstraints - Get constraints
type GetEntryConstraints interface {
	[]EntryV1
}

// GetEntry - generic get function with context and uri for resource selection and filtering
func GetEntry[T GetEntryConstraints](ctx any, uri string) (t T, status *runtime.Status) {
	defer log2.Log(ctx, "GET", uri, log2.NewStatusCodeClosure(&status))()
	var e runtime.LogError

	u, err := url.Parse(uri)
	if err != nil {
		status = runtime.NewStatusError(runtime.StatusInvalidContent, getEntryLoc2, err)
		e.Handle(status, runtime.RequestId(ctx), "")
		return
	}
	if runtime.IsDebugEnvironment() {
		if fn := http2.GetHandlerProxy(ctx); fn != nil {
			a, status1 := fn(ctx, uri, "")
			if !status1.OK() {
				e.Handle(status, runtime.RequestId(ctx), "")
				return t, status1
			}
			return entryFromAny[T](a)
		}
	}
	t, status = getEntry[T](ctx, u, "")
	if !status.OK() {
		e.Handle(status, runtime.RequestId(ctx), getEntryLoc2)
		return
	}
	return t, runtime.NewStatusOK()
}

// PostEntryConstraints - Post constraints
type PostEntryConstraints interface {
	[]EntryV1 | []byte | runtime.Nillable
}

// PostEntry - exchange function
func PostEntry[T PostEntryConstraints](ctx any, method, uri, variant string, body T) (any, *runtime.Status) {
	var e runtime.LogError

	req, status := http2.NewRequest(ctx, method, uri, variant)
	if !status.OK() {
		e.Handle(status, runtime.RequestId(ctx), postEntryLoc2)
		return nil, status
	}
	return postWrapper(ctx, req, body)
}

// HttpHandler - Http endpoint
func HttpHandler(w http.ResponseWriter, r *http.Request) {
	httpWrapper(nil, w, r)
}
