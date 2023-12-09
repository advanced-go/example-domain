package entryv1

import (
	"context"
	"fmt"
	"github.com/advanced-go/core/access"
	"github.com/advanced-go/core/http2"
	"github.com/advanced-go/core/http2/http2test"
	"github.com/advanced-go/core/runtime"
	"github.com/advanced-go/example-domain/timeseries/context2"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func _Example_HttpHandler() {
	access.EnableTestLogHandler()

	rec := http2.NewRecorder()
	//req, _ := http.NewRequest("", "https://localhost:8080/advanced-go/example-domain/timeseries/entry", nil)
	//req.Header.Add(http2.ContentLocation, EntryV1Variant)
	//HttpHandler(rec, req)
	fmt.Printf("test: HttpHandler() -> %v", rec.Code)

	//Output:
	//test: HttpHandler() -> 404

}

func Test_httpHandler(t *testing.T) {
	deleteEntries()
	//fmt.Printf("test: Start Entries -> %v\n", len(list))
	type args struct {
		req    string
		resp   string
		status runtime.Status
	}
	tests := []struct {
		name string
		args args
	}{
		{"put-entries", args{req: "put-req-v1.txt", resp: "put-resp-v1.txt"}},
		{"get-entries", args{req: "get-req-v1.txt", resp: "get-resp-v1.txt"}},
		//	{"get-entries-by-controller", args{req: "get-ctrl-req.txt", resp: "get-ctrl-resp.txt"}},
		{"delete-entries", args{req: "delete-req-v1.txt", resp: "delete-resp-v1.txt"}},
	}
	for _, tt := range tests {
		failures, req, resp := http2test.ReadHttp("file://[cwd]/entryv1test/resource/", tt.args.req, tt.args.resp)
		if failures != nil {
			t.Errorf("ReadHttp() failures = %v", failures)
			continue
		}
		var ctx context.Context
		if tt.args.status != nil {
			ctx = context2.NewStatusContext(nil, tt.args.status)
		}
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			// ignoring returned status as any errors will be reflected in the response StatusCode
			httpHandler[runtime.Output](ctx, w, req)

			// kludge for BUG in response recorder
			w.Result().Header = w.Header()

			// test status code
			if w.Result().StatusCode != resp.StatusCode {
				t.Errorf("StatusCode got = %v, want %v", w.Result().StatusCode, resp.StatusCode)
			} else {
				// test headers if needed - test2.Headers(w.Result(),resp,names... string) (failures []Args)

				// test content size and unmarshal types
				var gotT, wantT []Entry
				var content bool
				failures, content, gotT, wantT = http2test.Content[[]Entry](w.Result(), resp, testBytes)
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

func testBytes(got *http.Response, gotBytes []byte, want *http.Response, wantBytes []byte) []http2test.Args {
	//fmt.Printf("got = %v\n[len:%v]\n", string(gotBytes), len(gotBytes))
	//fmt.Printf("want = %v\n[len:%v]\n", string(wantBytes), len(wantBytes))
	return nil
}

func Errorf(t *testing.T, failures []http2test.Args) {
	for _, arg := range failures {
		t.Errorf("%v got = %v want = %v", arg.Item, arg.Got, arg.Want)
	}
}
