package google

import (
	"context"
	"fmt"
	"github.com/go-ai-agent/core/httpx"
	"github.com/go-ai-agent/core/log"
	"github.com/go-ai-agent/core/runtime/runtimetest"
	"net/http"
	"net/url"
)

func Example_TypeHandler() {
	ctx := log.NewAccessContext(context.Background())
	req, _ := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080"+pkgPath+"?q=test", nil)
	resp, status := Do(nil, req, nil)
	if buf, ok := resp.([]byte); ok {
		if buf != nil {
		}
	}

	fmt.Printf("test: TypeHandler(%v) -> [status:%v] [content-type:%v] [content-length:%v]\n", req.URL.String(), status, status.Header().Get(httpx.ContentType), status.Header().Get(httpx.ContentLength))

	//Output:
	//test: typeHandler(http://localhost:8080/go-ai-agent/example-domain/google/search?q=test) -> [status:OK] [content-type:text/html; charset=ISO-8859-1] [content-length:100835]

}

func Example_typeHandler() {
	req, _ := http.NewRequest("", "http://localhost:8080"+pkgPath+"?q=test", nil)
	resp, status := doHandler[runtimetest.DebugError](nil, req, nil)
	if buf, ok := resp.([]byte); ok {
		if buf != nil {
		}
	}

	fmt.Printf("test: typeHandler(%v) -> [status:%v] [content-type:%v] [content-length:%v]\n", req.URL.String(), status, status.Header().Get(httpx.ContentType), status.Header().Get(httpx.ContentLength))

	//Output:
	//test: typeHandler(http://localhost:8080/go-ai-agent/example-domain/google/search?q=test) -> [status:OK] [content-type:text/html; charset=ISO-8859-1] [content-length:100835]

}
func Example_httpHandler() {
	r := httpx.NewRecorder()

	req, _ := http.NewRequest("", "http://localhost:8080"+pkgPath+"?q=test", nil)
	status := httpHandler[runtimetest.DebugError](r, req)
	r.Result().Header = r.Header()
	buf, status1 := httpx.ReadAll(r.Result().Body)
	fmt.Printf("test: ReadAll() -> [status:%v] [body:%v]\n", status1, len(buf))

	fmt.Printf("test: httpHandler(%v) -> [status:%v] [content-type:%v] [content-length:%v]\n", req.URL.String(), status, r.Result().Header.Get(httpx.ContentType), r.Result().Header.Get(httpx.ContentLength))

	//Output:test: ReadAll() -> [status:OK] [body:100705]
	//test: httpHandler(http://localhost:8080/go-ai-agent/example-domain/google/search?q=test) -> [status:OK] [content-type:text/html; charset=utf-8] [content-length:100705]

}

func Example_Resolver() {
	// Resolve the content to a file
	fileUri := "file://[cwd]/resource/query-result.txt"
	httpx.AddResolver(func(s string) string {
		return fileUri
	},
	)
	u, _ := url.Parse(fileUri)
	buf, err := httpx.ReadFile(u)
	fmt.Printf("test: ReadFile() -> [err:%v] [buf:%v]\n", err, string(buf))

	req, _ := http.NewRequest("", pkgPath, nil)
	result, status := doHandler[runtimetest.DebugError](nil, req, nil)
	str := ""
	if buf1, ok := result.([]byte); ok {
		str = string(buf1)
	}
	fmt.Printf("test: typeHandler() [status:%v] [content:%v]\n", status, str)

	//Output:
	//test: ReadFile() -> [err:<nil>] [buf:This is an alternate result for a Google query.]
	//test: typeHandler() [status:OK] [content:This is an alternate result for a Google query.]

}
