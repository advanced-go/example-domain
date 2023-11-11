package activity

import (
	"fmt"
	"github.com/go-ai-agent/core/httpx/httpxtest"
	"github.com/go-ai-agent/core/runtime/runtimetest"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func Test_httpHandler(t *testing.T) {

	deleteEntries()
	fmt.Printf("test: Start Entries -> %v\n", len(list))
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
		{"get-entries-by-type", args{req: "get-type-req.txt", resp: "get-type-resp.txt"}},
		{"delete-entries", args{req: "delete-req.txt", resp: "delete-resp.txt"}},
	}
	for _, tt := range tests {
		failures, req, resp := httpxtest.ReadHttp("file://[cwd]/activitytest/resource/", tt.args.req, tt.args.resp)
		if failures != nil {
			t.Errorf("ReadHttp() failures = %v", failures)
			continue
		}
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			// ignoring returned status as any errors will be reflected in the response StatusCode
			httpHandler[runtimetest.DebugError](nil, w, req)

			// kludge for BUG in response recorder
			w.Result().Header = w.Header()

			// test status code
			if w.Result().StatusCode != resp.StatusCode {
				t.Errorf("StatusCode got = %v, want %v", w.Result().StatusCode, resp.StatusCode)
			} else {
				// test headers if needed - test2.Headers(w.Result(),resp,names... string) (failures []Args)

				// test content size and unmarshal types
				var gotT, wantT []EntryV1
				var content bool
				failures, content, gotT, wantT = httpxtest.Content[[]EntryV1](w.Result(), resp, testBytes)
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
	fmt.Printf("test: End Entries -> %v\n", len(list))
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

//t.Run(tt.name, func(t *testing.T) {
//	if got := entryHandler(tt.args.w, tt.args.r); !reflect.DeepEqual(got, tt.want) {
//		t.Errorf("entryHandler() = %v, want %v", got, tt.want)
//	}
//})
