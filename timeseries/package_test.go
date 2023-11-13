package timeseries

import (
	"fmt"
	"github.com/go-ai-agent/core/log2"
	"github.com/go-ai-agent/core/runtime"
	"time"
)

var entries = []EntryV1{{
	Traffic:   "ingress",
	Start:     time.Now().UTC(),
	Duration:  750,
	RequestId: "98765",
	Url:       "https://www.somestupiddomain.com/help",
}}

func getProxy(ctx any, uri, variant string) (any, *runtime.Status) {
	fmt.Printf("test: getProxy() -> in proxy\n")
	return entries, runtime.NewStatusOK() //http.StatusInternalServerError)
}

func Example_GetWithProxy() {
	log2.EnableDebugAccessHandler()
	ctx := runtime.NewRequestIdContext(nil, "123-request-456")
	ctx = runtime.NewProxyContext(ctx, getProxy)
	e, status := Get[[]EntryV1](ctx, "https://google.com/search")
	fmt.Printf("test: Get() -> [status:%v] %v", status, e)

	//Output:
	//test: getProxy() -> in proxy

}
