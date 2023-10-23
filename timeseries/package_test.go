package timeseries

import (
	"fmt"
	"github.com/go-ai-agent/core/httpx"
	"github.com/go-ai-agent/core/httpx/httpxtest"
	"reflect"

	//test2 "github.com/go-ai-agent/core/exchange/httptest"
	"github.com/go-ai-agent/core/runtime"
	"net/http/httptest"
	"testing"
)

func _Example_EntryHandler_PUT() {
	deleteEntries()
	prevCnt := len(list)

	s := "file://[cwd]/timeseriestest/resource/timeseries-entry-put-req.txt"
	req, _ := httpx.ReadRequest(runtime.ParseRaw(s))
	//fmt.Printf("test: ReadRequest(%v) [err:%v] %v\n", s, err, req)

	rec := httpx.NewRecorder()
	EntryHandler(rec, req)
	currCnt := len(list)
	//fmt.Printf("test: list %v\n", list)

	s1 := "file://[cwd]/timeseriestest/resource/timeseries-entry-put-resp.txt"
	target, _ := httpx.ReadResponse(runtime.ParseRaw(s1))
	//fmt.Printf("test: ReadResponse(%v) [err:%v] %v\n", s1, err1, target)

	src := rec.Result()
	fmt.Printf("test: EntryHandler() [prevCnt:%v] [currCnt:%v] [targetStatus:%v] [srcStatus:%v]\n", prevCnt, currCnt, target.StatusCode, src.StatusCode)

	//Output:
	//test: EntryHandler() [prevCnt:0] [currCnt:2] [targetStatus:200] [srcStatus:200]

}

func _Example_EntryHandler_GET() {
	fmt.Printf("test: list %v\n", list)

	s := "file://[cwd]/timeseriestest/resource/timeseries-entry-get-req.txt"
	req, err := httpx.ReadRequest(runtime.ParseRaw(s))
	fmt.Printf("test: ReadRequest(%v) [err:%v] %v\n", s, err, req)

	rec := httptest.NewRecorder()
	EntryHandler(rec, req)
	fmt.Printf("test: list %v\n", list)

	s1 := "file://[cwd]/timeseriestest/resource/timeseries-entry-get-resp.txt"
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

func TestEntryHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			EntryHandler(tt.args.w, tt.args.r)
		})
	}
}

func TestGetEntries(t *testing.T) {
	tests := []struct {
		name string
		want []Entry
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetEntries(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEntries() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetEntriesByController(t *testing.T) {
	type args struct {
		ctrl string
	}
	tests := []struct {
		name string
		args args
		want []Entry
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetEntriesByController(tt.args.ctrl); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEntriesByController() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMarshalEntry(t *testing.T) {
	type args struct {
		entry []Entry
	}
	tests := []struct {
		name  string
		args  args
		want  []byte
		want1 *runtime.Status
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := MarshalEntry(tt.args.entry)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalEntry() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("MarshalEntry() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestUnmarshalEntry(t *testing.T) {
	type args struct {
		buf []byte
	}
	tests := []struct {
		name  string
		args  args
		want  []Entry
		want1 *runtime.Status
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := UnmarshalEntry(tt.args.buf)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnmarshalEntry() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("UnmarshalEntry() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_deleteEntries(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deleteEntries()
		})
	}
}

*/

func Test_entryHandler(t *testing.T) {
	deleteEntries()
	type args struct {
		req  string
		resp string
	}
	tests := []struct {
		name string
		args args
	}{
		{"put entries", args{req: "timeseries-entry-put-req.txt", resp: "timeseries-entry-put-req.txt"}},
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
			entryHandler(w, req)

			// test status code
			if w.Result().StatusCode != resp.StatusCode {
				t.Errorf("StatusCode got = %v, want %v", w.Result().StatusCode, resp.StatusCode)
			}

			// test headers if needed - test2.Headers(w.Result(),resp,names... string) (failures []Args)

			// test content size and unmarshal types
			var gotT, wantT []Entry
			failures, gotT, wantT = httpxtest.Content[[]Entry](w.Result(), resp)
			if failures != nil {
				t.Errorf("Content() failures = %v", failures)
			}

			// test types
			if !reflect.DeepEqual(gotT, wantT) {
				t.Errorf("DeepEqual() got = %v, want %v", gotT, wantT)
			}
		})
	}
}
