package google

import (
	"fmt"
	"github.com/go-ai-agent/core/http2"
	"github.com/go-ai-agent/core/io2"
	"github.com/go-ai-agent/core/runtime/runtimetest"
	"net/http"
	"net/url"
)

func Example_httpHandler() {
	r := http2.NewRecorder()

	req, _ := http.NewRequest("", "http://localhost:8080"+pkgPath+"?q=test", nil)
	status := httpHandler[runtimetest.DebugError](r, req)
	r.Result().Header = r.Header()
	buf, status1 := io2.ReadAll(r.Result().Body)
	fmt.Printf("test: ReadAll() -> [status:%v] [body:%v]\n", status1, len(buf))

	fmt.Printf("test: httpHandler(%v) -> [status:%v] [content-type:%v] [content-length:%v]\n", req.URL.String(), status, r.Result().Header.Get(http2.ContentType), r.Result().Header.Get(http2.ContentLength))

	//Output:test: ReadAll() -> [status:OK] [body:100705]
	//test: httpHandler(http://localhost:8080/go-ai-agent/example-domain/google/search?q=test) -> [status:OK] [content-type:text/html; charset=utf-8] [content-length:100705]

}

func Example_Resolver() {
	// Resolve the content to a file
	fileUri := "file://[cwd]/resource/query-result.txt"
	http2.AddResolver(func(s string) string {
		return fileUri
	},
	)
	u, _ := url.Parse(fileUri)
	buf, err := io2.ReadFile(u)
	fmt.Printf("test: ReadFile() -> [err:%v] [buf:%v]\n", err, string(buf))

	req, _ := http.NewRequest("", pkgPath, nil)
	result, status := doHandler[runtimetest.DebugError](nil, req, nil)
	str := ""
	if buf1, ok := result.([]byte); ok {
		str = string(buf1)
	}
	fmt.Printf("test: doHandler() [status:%v] [content:%v]\n", status, str)

	//Output:
	//test: ReadFile() -> [err:<nil>] [buf:This is an alternate result for a Google query.]
	//test: doHandler() [status:OK] [content:This is an alternate result for a Google query.]

}
