package entryv2

import (
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx/httpxtest"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

const (
	emptyEntry = "file://[cwd]/entryv2test/resource/empty.json"
	validEntry = "file://[cwd]/entryv2test/resource/timeseries-entry-v2.json"
)

func _Example_HttpHandler() {
	//access.EnableTestLogger()

	rec := httptest.NewRecorder()
	//req, _ := http.NewRequest("", "https://localhost:8080/advanced-go/example-domain/timeseries/entry", nil)
	//req.Header.Add(http2.ContentLocation, EntryV1Variant)
	//HttpHandler(rec, req)
	fmt.Printf("test: HttpHandler() -> %v", rec.Code)

	//Output:
	//test: HttpHandler() -> 404

}

func Test_httpHandler(t *testing.T) {
	deleteEntries(nil)
	//fmt.Printf("test: Start Entries -> %v\n", len(list))
	type args struct {
		req    string
		resp   string
		result any
	}
	tests := []struct {
		name string
		args args
	}{
		{"put-entries", args{req: "put-req-v2.txt", resp: "put-resp-v2.txt", result: map[string]string{"addEntries": ""}}},
		{"get-entries", args{req: "get-req-v2.txt", resp: "get-resp-v2.txt", result: map[string]string{"getEntries": validEntry}}},
		//	{"get-entries-by-controller", args{req: "get-ctrl-req.txt", resp: "get-ctrl-resp.txt",resultemptyEntry}},
		{"delete-entries", args{req: "delete-req-v2.txt", resp: "delete-resp-v2.txt", result: map[string]string{"deleteEntries": ""}}},
	}
	for _, tt := range tests {
		failures, req, resp := httpxtest.ReadHttp("file://[cwd]/entryv2test/resource/", tt.args.req, tt.args.resp)
		if failures != nil {
			t.Errorf("ReadHttp() failures = %v", failures)
			continue
		}
		lookup.SetOverride(tt.args.result)
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			// ignoring returned status as any errors will be reflected in the response StatusCode
			httpHandler[core.Output](w, req)

			// test status code
			if w.Result().StatusCode != resp.StatusCode {
				t.Errorf("StatusCode got = %v, want %v", w.Result().StatusCode, resp.StatusCode)
			} else {
				// test headers if needed - test2.Headers(w.Result(),resp,names... string) (failures []Args)

				// test content size and unmarshal types
				var gotT, wantT []Entry
				var content bool
				failures, content, gotT, wantT = httpxtest.Content[[]Entry](w.Result(), resp, testBytes)
				if failures != nil {
					//t.Errorf("Content() failures = %v", failures)
					Errorf(t, failures)
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
	//fmt.Printf("test: End Entries -> %v\n", len(listV2))
}

func testBytes(got *http.Response, gotBytes []byte, want *http.Response, wantBytes []byte) []httpxtest.Args {
	//fmt.Printf("got = %v\n[len:%v]\n", string(gotBytes), len(gotBytes))
	//fmt.Printf("want = %v\n[len:%v]\n", string(wantBytes), len(wantBytes))
	return nil
}

func Errorf(t *testing.T, failures []httpxtest.Args) {
	for _, arg := range failures {
		t.Errorf("%v got = %v want = %v", arg.Item, arg.Got, arg.Want)
	}
}
