package google

import (
	"fmt"
	"github.com/go-ai-agent/core/httpx"
	"github.com/go-ai-agent/core/runtime/runtimetest"
	"net/http"
)

func Example_DoHandler() {

}

// Example_searchHandler_ConnectivityError - this resolves the host to http://localhost:8080, so it will fail
func _Example_searchHandler_ConnectivityError() {
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
	//[[] github.com/go-ai-agent/core/exchange/Do [Get "http://localhost:8080/search?q=test": dial tcp [::1]:8080: connectex: No connection could be made because the target machine actively refused it.]]
	//[[] github.com/go-ai-agent/core/exchange/Do [Get "http://localhost:8080/search?q=test": dial tcp [::1]:8080: connectex: No connection could be made because the target machine actively refused it.]]
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

	// routing to localhost, so should see connectivity errors
	w := httpx.NewRecorder()
	searchHandler[runtimetest.DebugError](w, req)
	w.Result().Header = w.Header()
	buf, status := httpx.ReadAll(w.Result().Body)
	fmt.Printf("test: Response() [status:%v] [content:%v]\n", status, len(buf) > 0)

	//Output:
	//test: NewRequest(https://github.com/go-ai-agent/example-domain/google/search?q=test) [err:<nil>] [req:true]
	//test: Response() [status:OK] [content:true]

}
