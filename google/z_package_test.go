package google

import (
	"fmt"
	"github.com/go-ai-agent/core/exchange"
	"github.com/go-ai-agent/core/httpx"
	"github.com/go-ai-agent/core/runtime/runtimetest"
	"net/http"
)

func _Example_DoHandler() {
	req, _ := http.NewRequest("", "http://localhost:8080/"+SearchEndpoint+"?q=test", nil)
	resp, err := DoHandler(req)
	fmt.Printf("test: DoHandler(%v) -> [err:%v] [status:%v] [content-length:%v]\n", req.URL.String(), err, resp.StatusCode, resp.Header.Get(httpx.ContentLength))

	httpx.AddResolver(func(s string) string {
		return "https://www.google.com" + s
	})

	req, _ = http.NewRequest("", "https://www.google.com/search?q=test", nil)
	resp, err = DoHandler(req)
	fmt.Printf("test: DoHandler(%v) -> [err:%v] [status:%v] [content-length:%v]\n", req.URL.String(), err, resp.StatusCode, resp.Header.Get(httpx.ContentLength))

	//Output:
	//test: DoHandler(http://localhost:8080//go-ai-agent/example-domain/google/search?q=test) -> [err:Get "http://localhost:8080/search?q=test": dial tcp [::1]:8080: connectex: No connection could be made because the target machine actively refused it.] [status:500] [content-length:]
	//test: DoHandler(https://www.google.com/search?q=test) -> [err:<nil>] [status:200] [content-length:100656]

}

// Example_searchHandler_ConnectivityError - this resolves the host to http://localhost:8080, so it will fail
func Example_searchHandler_ConnectivityError() {
	uri := "https://github.com" + SearchEndpoint + "?q=test"
	req, err := http.NewRequest("", uri, nil)
	fmt.Printf("test: NewRequest(%v) [err:%v] [req:%v]\n", uri, err, req != nil)

	// routing to localhost, so should see connectivity errors
	w := httpx.NewRecorder()
	searchHandler[runtimetest.DebugError](w, req)
	w.Result().Header = w.Header()
	// nil the body so test will succeed
	w.Result().Body = nil
	fmt.Printf("test: Response() %v\n", w.Result())

	//Output:
	//test: NewRequest(https://github.com/go-ai-agent/example-domain/google/search?q=test) [err:<nil>] [req:true]
	//{ "id":null, "l":"github.com/go-ai-agent/core/exchange/Do", "o":"github.com/go-ai-agent/example-domain/google/searchHandler" "err" : [ "Get "http://localhost:8080/search?q=test": dial tcp [::1]:8080: connectex: No connection could be made because the target machine actively refused it." ] }
	//test: Response() &{500 Internal Server Error 500 HTTP/1.1 1 1 map[] <nil> -1 [] false false map[] <nil> <nil>}

}

// Example_searchHandler_OK - this configures a new resolver, to https://www.google.com, so this should work
func Example_searchHandler_OK() {
	httpx.AddResolver(func(s string) string {
		return "https://www.google.com" + s
	})

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
