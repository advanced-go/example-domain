package timeseries

import (
	"fmt"
	"github.com/go-ai-agent/core/exchange"
	"github.com/go-ai-agent/core/runtime"
	"net/http/httptest"
)

func Example_EntryHandler_PUT() {
	deleteEntries()
	prevCnt := len(list)

	s := "file://[cwd]/timeseriestest/resource/timeseries-entry-put-req.txt"
	req, _ := exchange.ReadRequest(runtime.ParseRaw(s))
	//fmt.Printf("test: ReadRequest(%v) [err:%v] %v\n", s, err, req)

	rec := httptest.NewRecorder()
	EntryHandler(rec, req)
	currCnt := len(list)
	//fmt.Printf("test: list %v\n", list)

	s1 := "file://[cwd]/timeseriestest/resource/timeseries-entry-put-resp.txt"
	target, _ := exchange.ReadResponse(runtime.ParseRaw(s1))
	//fmt.Printf("test: ReadResponse(%v) [err:%v] %v\n", s1, err1, target)

	src := rec.Result()
	fmt.Printf("test: EntryHandler() [prevCnt:%v] [currCnt:%v] [targetStatus:%v] [srcStatus:%v]\n", prevCnt, currCnt, target.StatusCode, src.StatusCode)

	//Output:
	//test: EntryHandler() [prevCnt:0] [currCnt:2] [targetStatus:200] [srcStatus:200]

}

func Example_EntryHandler_GET() {
	fmt.Printf("test: list %v\n", list)

	s := "file://[cwd]/timeseriestest/resource/timeseries-entry-get-req.txt"
	req, err := exchange.ReadRequest(runtime.ParseRaw(s))
	fmt.Printf("test: ReadRequest(%v) [err:%v] %v\n", s, err, req)

	rec := httptest.NewRecorder()
	EntryHandler(rec, req)
	fmt.Printf("test: list %v\n", list)

	s1 := "file://[cwd]/timeseriestest/resource/timeseries-entry-get-resp.txt"
	target, err1 := exchange.ReadResponse(runtime.ParseRaw(s1))
	fmt.Printf("test: ReadResponse(%v) [err:%v] %v\n", s1, err1, target)

	src := rec.Result()
	fmt.Printf("test: Response() %v\n", src)

	//Output:
}
