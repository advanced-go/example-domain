package timeseries

import (
	"fmt"
	"github.com/advanced-go/core/http2"
	"github.com/advanced-go/core/runtime"
	"github.com/advanced-go/core/runtime/runtimetest"
	"net/http"
	"net/url"
	"reflect"
	"time"
)

var entries = []EntryV1{{
	Traffic:   "ingress",
	Start:     time.Now().UTC(),
	Duration:  750,
	RequestId: "98765",
	Url:       "https://www.somestupiddomain.com/help",
}}

func Example_PkgUri() {
	pkgUri2 := reflect.TypeOf(any(pkg{})).PkgPath()
	pkgPath2 := runtime.PathFromUri(pkgUri2)
	entryV1 := pkgUri2 + "/" + reflect.TypeOf(EntryV1{}).Name()
	entryV2 := pkgUri2 + "/" + reflect.TypeOf(EntryV2{}).Name()

	fmt.Printf("test: PkgUri         = \"%v\"\n", pkgUri2)
	fmt.Printf("test: PkgPath        = \"%v\"\n", pkgPath2)
	fmt.Printf("test: Pattern        = \"%v\"\n", pkgPath2+"/")
	fmt.Printf("test: EntryV1Variant = \"%v\"\n", entryV1)
	fmt.Printf("test: EntryV2Variant = \"%v\"\n", entryV2)

	//Output:
	//test: PkgUri         = "github.com/advanced-go/example-domain/timeseries"
	//test: PkgPath        = "/advanced-go/example-domain/timeseries"
	//test: Pattern        = "/advanced-go/example-domain/timeseries/"
	//test: EntryV1Variant = "github.com/advanced-go/example-domain/timeseries/EntryV1"
	//test: EntryV2Variant = "github.com/advanced-go/example-domain/timeseries/EntryV2"

}

func getProxy(h http.Header, uri *url.URL) (any, runtime.Status) {
	fmt.Printf("test: getProxy() -> in proxy\n")
	return entries, runtime.NewStatusOK() //http.StatusInternalServerError)
}

func _Example_GetWithProxy() {
	//ctx := runtime.NewRequestIdContext(nil, "get-654-321")
	//ctx = runtime.NewProxyContext(ctx, getProxy)
	h := make(http.Header)
	u, _ := url.Parse("https://google.com/search")
	e, status := getEntryHandler[[]EntryV1](nil, h, u)
	fmt.Printf("test: getEntryHandler[[]EntryV1]() -> [status:%v] [entries:%v]\n", status, len(e))

	//Output:
	//test: getProxy() -> in proxy
	//test: getEntryHandler[[]EntryV1]() -> [status:OK] [entries:1]

}

func postProxy(r *http.Request, body any) (any, runtime.Status) {
	fmt.Printf("test: postProxy() -> in proxy\n")
	return nil, runtime.NewStatus(http.StatusServiceUnavailable)
}

func _Example_PostWithProxy() {
	ctx := runtime.NewRequestIdContext(nil, "post-123-456")
	ctx = runtime.NewProxyContext(ctx, postProxy)
	req, _ := http2.NewRequest(ctx, "PUT", "https://google.com/search", EntryV1Variant, nil)
	e, status := postEntryHandler(ctx, req, nil)
	fmt.Printf("test: postEntryHandler[runtimetest.DebugError]() -> [status:%v] %v\n", status, e)

	//Output:
	//test: postProxy() -> in proxy
	//test: postEntryHandler[runtimetest.DebugError]() -> [status:Service Unavailable] <nil>

}

func httpProxy(w http.ResponseWriter, r *http.Request) runtime.Status {
	fmt.Printf("test: httpProxy() -> in proxy\n")
	return runtime.NewStatus(http.StatusGatewayTimeout)
}

func _Example_HttpWithProxy() {
	ctx := runtime.NewRequestIdContext(nil, "http-456-789")
	ctx = runtime.NewProxyContext(ctx, httpProxy)
	rec := http2.NewRecorder()
	req, _ := http.NewRequestWithContext(ctx, "DELETE", "https://www.google.com/search", nil)
	req.Header.Add(http2.ContentLocation, EntryV1Variant)
	status := httpHandler[runtimetest.DebugError](ctx, rec, req)
	fmt.Printf("test: httpHandler() -> [status:%v]\n", status)

	//Output:
	//test: httpProxy() -> in proxy
	//test: httpHandler() -> [status:Timeout]

}
