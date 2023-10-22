package timeseries

import (
	"fmt"
	"github.com/go-ai-agent/core/exchange"
	"github.com/go-ai-agent/core/runtime"
	"net/http/httptest"
)

func Example_EntryHandler_PUT() {
	s := "file://[cwd]/timeseriestest/resource/timeseries-entry-put-req.txt"
	req, err := exchange.ReadRequest(runtime.ParseRaw(s))
	fmt.Printf("test: ReadRequest(%v) [err:%v] %v\n", s, err, req)

	rec := httptest.NewRecorder()
	EntryHandler(rec, req)

	s1 := "file://[cwd]/timeseriestest/resource/timeseries-entry-put-resp.txt"
	target, err1 := exchange.ReadResponse(runtime.ParseRaw(s1))
	fmt.Printf("test: ReadResponse(%v) [err:%v] %v\n", s1, err1, target)

	src := rec.Result()
	fmt.Printf("test: Response() %v\n", src)

	//Output:
}
