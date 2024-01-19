package service

import (
	"fmt"
	"github.com/advanced-go/core/access"
	"github.com/advanced-go/core/runtime"
	"net/http"
	"net/http/httptest"
)

func ExampleSearchHandler() {
	access.EnableTestLogger()

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("", "http://localhost:8080/github/advanced-go/example-domain/service:search?q=golang", nil)
	status := searchHandler[runtime.Output](rec, req)
	resp := rec.Result()
	buf, status0 := runtime.ReadAll(resp.Body, nil)
	fmt.Printf("test: searchHandler() -> [code:%v] [read-status:%v] [status:%v] [content:%v]\n", rec.Result().StatusCode, status0, status, buf != nil && len(buf) > 0)

	//Output:
	//test: searchHandler() -> [code:500] [read-status:OK] [status:Internal Error [Get "http://localhost:8081/github/advanced-go/search/provider:search?q=golang": dial tcp [::1]:8081: connectex: No connection could be made because the target machine actively refused it.]] [content:false]

}
