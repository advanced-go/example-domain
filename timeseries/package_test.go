package timeseries

import (
	"fmt"
	"github.com/go-ai-agent/core/httpx"
	"github.com/go-ai-agent/core/httpx/httpxtest"
	"github.com/go-ai-agent/core/runtime"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func Example_PkgUri() {
	fmt.Printf("test: PkgUrl %v\n", PkgUrl)
	fmt.Printf("test: PkgUri %v\n", PkgUri)
	fmt.Printf("test: EntryPath %v\n", EntryPath)

	//var e []entry
	//fmt.Printf("test: e [len:%v] [nil:%v]\n ", len(e), e == nil)

	//Output:
	//test: PkgUrl file://github.com/go-ai-agent/example-domain/timeseries
	//test: PkgUri github.com/go-ai-agent/example-domain/timeseries
	//test: EntryPath /go-ai-agent/example-domain/timeseries/entry

}

func _Example_EntryHandler_PUT() {
	deleteEntries()
	prevCnt := len(list)

	s := "file://[cwd]/timeseriestest/resource/put-req.txt"
	req, _ := httpx.ReadRequest(runtime.ParseRaw(s))
	//fmt.Printf("test: ReadRequest(%v) [err:%v] %v\n", s, err, req)

	rec := httpx.NewRecorder()
	EntryHandler(rec, req)
	currCnt := len(list)
	//fmt.Printf("test: list %v\n", list)

	s1 := "file://[cwd]/timeseriestest/resource/put-resp.txt"
	target, _ := httpx.ReadResponse(runtime.ParseRaw(s1))
	//fmt.Printf("test: ReadResponse(%v) [err:%v] %v\n", s1, err1, target)

	src := rec.Result()
	fmt.Printf("test: EntryHandler() [prevCnt:%v] [currCnt:%v] [targetStatus:%v] [srcStatus:%v]\n", prevCnt, currCnt, target.StatusCode, src.StatusCode)

	//Output:
	//test: EntryHandler() [prevCnt:0] [currCnt:2] [targetStatus:200] [srcStatus:200]

}

func _Example_EntryHandler_GET() {
	fmt.Printf("test: list %v\n", list)

	s := "file://[cwd]/timeseriestest/resource/get-req.txt"
	req, err := httpx.ReadRequest(runtime.ParseRaw(s))
	fmt.Printf("test: ReadRequest(%v) [err:%v] %v\n", s, err, req)

	rec := httptest.NewRecorder()
	EntryHandler(rec, req)
	fmt.Printf("test: list %v\n", list)

	s1 := "file://[cwd]/timeseriestest/resource/get-resp.txt"
	target, err1 := httpx.ReadResponse(runtime.ParseRaw(s1))
	fmt.Printf("test: ReadResponse(%v) [err:%v] %v\n", s1, err1, target)

	src := rec.Result()
	fmt.Printf("test: Response() %v\n", src)

	//Output:
}

/*
func TestAddEntry(t *testing.T) {
	type args struct {
		e []Entry
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AddEntry(tt.args.e)
		})
	}
}

func TestDoHandler(t *testing.T) {
	type args struct {
		req *http.Request
	}
	tests := []struct {
		name    string
		args    args
		want    *http.Response
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DoHandler(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("DoHandler() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DoHandler() got = %v, want %v", got, tt.want)
			}
		})
	}
}

*/

func Test_entryHandler(t *testing.T) {
	deleteEntries()
	//fmt.Printf("test: Start Entries -> %v\n", len(list))
	type args struct {
		req  string
		resp string
	}
	tests := []struct {
		name string
		args args
	}{
		{"put-entries", args{req: "put-req.txt", resp: "put-resp.txt"}},
		{"get-entries", args{req: "get-req.txt", resp: "get-resp.txt"}},
		{"get-entries-by-controller", args{req: "get-ctrl-req.txt", resp: "get-ctrl-resp.txt"}},
		{"delete-entries", args{req: "delete-req.txt", resp: "delete-resp.txt"}},
	}
	for _, tt := range tests {
		failures, req, resp := httpxtest.ReadHttp("file://[cwd]/timeseriestest/resource/", tt.args.req, tt.args.resp)
		if failures != nil {
			t.Errorf("ReadHttp() failures = %v", failures)
			continue
		}
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			// ignoring returned status as any errors will be reflected in the response StatusCode
			entryHandler[runtime.BypassError](w, req)

			// kludge for BUG in response recorder
			w.Result().Header = w.Header()

			// test status code
			if w.Result().StatusCode != resp.StatusCode {
				t.Errorf("StatusCode got = %v, want %v", w.Result().StatusCode, resp.StatusCode)
			} else {
				// test headers if needed - test2.Headers(w.Result(),resp,names... string) (failures []Args)

				// test content size and unmarshal types
				var gotT, wantT []entry
				var content bool
				failures, content, gotT, wantT = httpxtest.Content[[]entry](w.Result(), resp, testBytes)
				if failures != nil {
					t.Errorf("Content() failures = %v", failures)
				} else {
					// compare types
					if content {
						if !reflect.DeepEqual(gotT, wantT) {
							t.Errorf("DeepEqual() got = %v, want %v", gotT, wantT)
						}
					}
				}
			}
		})
	}
	fmt.Printf("test: End Entries -> %v\n", len(list))
}

func testBytes(got *http.Response, gotBytes []byte, want *http.Response, wantBytes []byte) []httpxtest.Args {
	//fmt.Printf("got = %v\n[len:%v]\n", string(gotBytes), len(gotBytes))
	//fmt.Printf("want = %v\n[len:%v]\n", string(wantBytes), len(wantBytes))
	return nil
}
