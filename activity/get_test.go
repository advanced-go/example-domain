package activity

import (
	"errors"
	"fmt"
	"github.com/advanced-go/core/http2"
	"github.com/advanced-go/core/io2"
	"github.com/advanced-go/core/json2"
	"github.com/advanced-go/core/runtime"
	"github.com/advanced-go/core/runtime/runtimetest"
	"net/http"
	"net/url"
)

func getProxy(h http.Header, uri *url.URL) (any, runtime.Status) {
	content := h.Get(http2.ContentLocation)
	if len(content) == 0 {
		return nil, runtime.NewStatusError(runtime.StatusInvalidContent, "getProxy", errors.New("content-location is empty"))
	}
	u, _ := url.Parse(content)
	buf, err := io2.ReadFile(u)
	if err != nil {
		return nil, runtime.NewStatusError(runtime.StatusInvalidContent, "getProxy", err)
	}
	var entries []EntryV1
	status := json2.Unmarshal(buf, &entries)
	return entries, status
}

func Example_getEntryHandler() {
	ctx := runtime.NewProxyContext(nil, getProxy)
	h := make(http.Header)
	h.Add(http2.ContentLocation, "file://[cwd]/activitytest/resource/activity.json")
	u, _ := url.Parse("http://advanced-go/example-domain/activity/entry")

	entries, status := getEntryHandler[[]EntryV1, runtimetest.DebugError](ctx, h, u)
	fmt.Printf("test: getEntryHandler() -> [entries:%v] [status:%v]\n", entries, status)

	//Output:
	//test

}
