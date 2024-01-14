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
	req, _ := http.NewRequest("", "http://localhost:8080/github.com/advanced-go/example-domain/service:search?q=golang", nil)
	searchHandler[runtime.Output](rec, req)
	resp := rec.Result()
	buf, status := runtime.NewBytes(resp)
	fmt.Printf("test: searchHandler() -> [code:%v] [status:%v] [data:%v]\n", rec.Code, status, string(buf))

	//Output:
	//test: searchHandler() -> [code:200] [status:OK] [data:this is a search result]

}
