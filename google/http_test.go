package google

import (
	"fmt"
	"github.com/advanced-go/core/http2"
	"github.com/advanced-go/core/runtime"
	"net/http"
	"net/http/httptest"
)

func Example_httpHandler() {
	r := httptest.NewRecorder()

	req, _ := http.NewRequest("", "http://localhost:8080"+"/"+PkgPath+"/search?q=test", nil)
	status := httpHandler[runtime.Output](r, req)
	r.Result().Header = r.Header()
	buf, status1 := runtime.NewBytes(r.Result())
	fmt.Printf("test: NewBytes() -> [status:%v] [body:%v]\n", status1, len(buf))

	fmt.Printf("test: httpHandler(%v) -> [status:%v] [content-type:%v] [content-length:%v]\n", req.URL.String(), status, r.Result().Header.Get(http2.ContentType), r.Result().Header.Get(http2.ContentLength))

	//Output:test: NewBytes() -> [status:OK] [body:100705]
	//test: httpHandler(http://localhost:8080/github.com/advanced-go/example-domain/google/search?q=test) -> [status:OK] [content-type:text/html; charset=utf-8] [content-length:100705]

}
