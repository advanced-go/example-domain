package google

import (
	"fmt"
	"github.com/advanced-go/core/http2"
	"github.com/advanced-go/core/http2/http2test"
	"github.com/advanced-go/core/io2"
	"github.com/advanced-go/core/runtime"
	"net/http"
	"net/url"
)

func Example_httpHandler() {
	r := http2test.NewRecorder()

	req, _ := http.NewRequest("", "http://localhost:8080"+"/"+PkgPath+"?q=test", nil)
	status := httpHandler[runtime.Output](r, req)
	r.Result().Header = r.Header()
	buf, status1 := io2.ReadAll(r.Result().Body)
	fmt.Printf("test: ReadAll() -> [status:%v] [body:%v]\n", status1, len(buf))

	fmt.Printf("test: httpHandler(%v) -> [status:%v] [content-type:%v] [content-length:%v]\n", req.URL.String(), status, r.Result().Header.Get(http2.ContentType), r.Result().Header.Get(http2.ContentLength))

	//Output:test: ReadAll() -> [status:OK] [body:100705]
	//test: httpHandler(http://localhost:8080/github.com/advanced-go/example-domain/google/search?q=test) -> [status:OK] [content-type:text/html; charset=utf-8] [content-length:100705]

}

func Example_Resolver() {
	// Resolve the content to a file
	fileUri := "file://[cwd]/resource/query-result.txt"
	addResolver(func(s string) string {
		return fileUri
	},
	)
	u, _ := url.Parse(fileUri)
	buf, err := io2.ReadFile(u)
	fmt.Printf("test: ReadFile() -> [err:%v] [buf:%v]\n", err, string(buf))

	req, _ := http.NewRequest("", PkgPath, nil)
	result, status := getHandler[runtime.Output](req)
	str := ""
	if buf1, ok := result.([]byte); ok {
		str = string(buf1)
	}
	fmt.Printf("test: getHandler() [status:%v] [content:%v]\n", status, str)

	//Output:
	//test: ReadFile() -> [err:<nil>] [buf:This is an alternate result for a Google query.]
	//test: getHandler() [status:OK] [content:This is an alternate result for a Google query.]

}
