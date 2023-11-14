package timeseries

import (
	"fmt"
	"github.com/go-ai-agent/core/http2"
	"github.com/go-ai-agent/core/log2"
	"github.com/go-ai-agent/core/runtime"
	"github.com/go-ai-agent/core/runtime/runtimetest"
	"net/http"
	"time"
)

var entries = []EntryV1{{
	Traffic:   "ingress",
	Start:     time.Now().UTC(),
	Duration:  750,
	RequestId: "98765",
	Url:       "https://www.somestupiddomain.com/help",
}}

func init() {
	log2.EnableDebugAccessHandler()
}

func getProxy(ctx any, uri, variant string) (any, *runtime.Status) {
	fmt.Printf("test: getProxy() -> in proxy\n")
	return entries, runtime.NewStatusOK() //http.StatusInternalServerError)
}

func Example_GetWithProxy() {
	ctx := runtime.NewRequestIdContext(nil, "get-654-321")
	ctx = runtime.NewProxyContext(ctx, getProxy)
	e, status := getEntryHandler[[]EntryV1, runtimetest.DebugError](ctx, "https://google.com/search")
	fmt.Printf("test: getEntryHandler[[]EntryV1]() -> [status:%v] [entries:%v]\n", status, len(e))

	//Output:
	//test: getProxy() -> in proxy
	//test: getEntryHandler[[]EntryV1]() -> [status:OK] [entries:1]

}

func postProxy(ctx any, r *http.Request, body any) (any, *runtime.Status) {
	fmt.Printf("test: postProxy() -> in proxy\n")
	return nil, runtime.NewStatus(http.StatusServiceUnavailable)
}

func Example_PostWithProxy() {
	ctx := runtime.NewRequestIdContext(nil, "post-123-456")
	ctx = runtime.NewProxyContext(ctx, postProxy)
	req, _ := http2.NewRequest(ctx, "PUT", "https://google.com/search", EntryV1Variant, nil)
	e, status := postEntryHandler[runtimetest.DebugError](ctx, req, nil)
	fmt.Printf("test: postEntryHandler[runtimetest.DebugError]() -> [status:%v] %v\n", status, e)

	//Output:
	//test: postProxy() -> in proxy
	//test: postEntryHandler[runtimetest.DebugError]() -> [status:Service Unavailable] <nil>

}

func httpProxy(ctx any, w http.ResponseWriter, r *http.Request) *runtime.Status {
	fmt.Printf("test: httpProxy() -> in proxy\n")
	return runtime.NewStatus(http.StatusGatewayTimeout)
}

func Example_HttpWithProxy() {
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
