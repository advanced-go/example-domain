package service

import (
	"fmt"
	"github.com/advanced-go/core/access"
	"github.com/advanced-go/core/runtime"
	"net/http"
	"net/http/httptest"
)

func _ExampleSearchHandler() {
	access.EnableTestLogger()

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("", "http://localhost:8080/github/advanced-go/example-domain/service:search?q=golang", nil)
	searchHandler[runtime.Output](rec, req)
	resp := rec.Result()
	buf, status := runtime.ReadAll(resp.Body, nil)
	fmt.Printf("test: searchHandler() -> [code:%v] [status:%v] [data:%v]\n", rec.Code, status, string(buf))

	//Output:
	//test: searchHandler() -> [code:200] [status:OK] [data:this is a search result]

}
