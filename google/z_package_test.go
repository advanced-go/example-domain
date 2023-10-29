package google

import (
	"fmt"
	"github.com/go-ai-agent/core/exchange"
	"github.com/go-ai-agent/core/httpx"
	"github.com/go-ai-agent/core/runtime/runtimetest"
	"net/http"
	"net/url"
)

func Example_DoHandler() {
	req, _ := http.NewRequest("", "http://localhost:8080"+SearchEndpoint+"?q=test", nil)
	resp, err := DoHandler(req)
	fmt.Printf("test: DoHandler(%v) -> [err:%v] [status:%v] [content-length:%v]\n", req.URL.String(), err, resp.StatusCode, resp.Header.Get(httpx.ContentLength))

	//Output:
	//test: DoHandler(http://localhost:8080/go-ai-agent/example-domain/google/search?q=test) -> [err:<nil>] [status:200] [content-length:108269]

}

// Example_searchHandler - this should work
func Example_searchHandler() {
	uri := "https://github.com" + SearchEndpoint + "?q=test"
	req, err := http.NewRequest("", uri, nil)
	fmt.Printf("test: NewRequest(%v) [err:%v] [req:%v]\n", uri, err, req != nil)

	// routing to https://www.google.com, so should work
	w := httpx.NewRecorder()
	searchHandler[runtimetest.DebugError](w, req)
	w.Result().Header = w.Header()
	buf, status := httpx.ReadAll(w.Result().Body)
	fmt.Printf("test: Response() [status:%v] [content:%v]\n", status, len(buf) > 0)

	//Output:
	//test: NewRequest(https://github.com/go-ai-agent/example-domain/google/search?q=test) [err:<nil>] [req:true]
	//test: Response() [status:OK] [content:true]

}

func Example_Resolver() {
	fileUri := "file://[cwd]/resource/query-result.txt"
	httpx.AddResolver(func(s string) string {
		return fileUri
	},
	)
	u, _ := url.Parse(fileUri)
	buf, err := httpx.ReadFile(u)
	fmt.Printf("test: ReadFile() -> [err:%v] [buf:%v]\n", err, string(buf))

	w := httpx.NewRecorder()
	req, _ := http.NewRequest("", pkgPath, nil)
	searchHandler[runtimetest.DebugError](w, req)
	w.Result().Header = w.Header()
	buf2, status := httpx.ReadAll(w.Result().Body)
	fmt.Printf("test: Response() [status:%v] [content:%v]\n", status, string(buf2))

	//Output:
	//test: ReadFile() -> [err:<nil>] [buf:This is an alternate result for a Google query.]
	//test: Response() [status:OK] [content:This is an alternate result for a Google query.]

}

func _Example_Proxy() {
	url := "http://localhost:8080" + SearchEndpoint + "?q=test"
	req, err := http.NewRequest("", url, nil)

	if err != nil {
		fmt.Printf("test: NewRequest(%v) -> [err:%v]\n", url, err)
	}
	resp, status := exchange.Do(req)
	fmt.Printf("test: Do(%v) -> [err:%v] [status:%v] [content-length:%v]\n", req.URL.String(), status, resp.StatusCode, resp.Header.Get(httpx.ContentLength))

	exchange.AddProxy(exchange.Proxy{Select: func(req *http.Request) bool { return true }, Do: do})
	req, err = http.NewRequest("", "https://www.google.com/search?q=test", nil)
	resp, status = exchange.Do(req)
	fmt.Printf("test: Do(%v) -> [err:%v] [status:%v] [content-length:%v]\n", req.URL.String(), status, resp.StatusCode, resp.Header.Get(httpx.ContentLength))

	//Output:
	//test: Do(http://localhost:8080/go-ai-agent/example-domain/google/search?q=test) -> [err:Internal Error [Get "http://localhost:8080/go-ai-agent/example-domain/google/search?q=test": dial tcp [::1]:8080: connectex: No connection could be made because the target machine actively refused it.]] [status:500] [content-length:]

}

func do(req *http.Request) (*http.Response, error) {
	//url, _ := url.Parse("https://www.google.com/search?q=test")
	//req.URL = url
	//url := "https://www.google.com/search?q=test"
	//req, _ = http.NewRequest("", url, nil)
	return exchange.Client.Do(req)
}
