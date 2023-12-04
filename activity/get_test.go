package activity

import (
	"errors"
	"fmt"
	"github.com/advanced-go/core/io2"
	"github.com/advanced-go/core/json2"
	"github.com/advanced-go/core/runtime"
	"net/http"
	"net/url"
)

func getProxy(h http.Header, uri *url.URL) (any, runtime.Status) {
	content := h.Get(ContentLocation)
	if len(content) == 0 {
		return nil, runtime.NewStatusError(runtime.StatusInvalidContent, "getProxy", errors.New("content-location is empty"))
	}
	u, _ := url.Parse(content)
	buf, err := io2.ReadFile(u)
	if err != nil {
		return nil, runtime.NewStatusError(runtime.StatusInvalidContent, "getProxy", err)
	}
	var entries []Entry
	status := json2.Unmarshal(buf, &entries)
	return entries, status
}

func _Example_getEntryHandler() {
	ctx := runtime.NewProxyContext(nil, getProxy)
	h := make(http.Header)
	h.Add(ContentLocation, "file://[cwd]/activitytest/resource/activity-entry-v1.json")
	u, _ := url.Parse("http://advanced-go/example-domain/activity/entry")

	entries, status := getEntryHandler[runtime.Output](ctx, h, u)
	fmt.Printf("test: getEntryHandler() -> [entries:%v] [status:%v]\n", entries, status)

	//Output:
	//test

}

func Example_getEntryFromPath() {
	location := "file://[cwd]/activitytest/resource/activity-entry-v1.json"

	//buf, status := getEntryFromPath(location)
	//fmt.Printf("test: getEntryFromPath() -> [buf:%v] [status:%v]\n", len(buf), status)

	entries, status2 := getEntryFromPath(location)
	fmt.Printf("test: getEntryFromPath() -> [entries:%v] [status:%v]\n", len(entries), status2)

	//Output:
	//test: getEntryFromPath() -> [entries:2] [status:OK]

}
