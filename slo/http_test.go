package slo

import (
	"github.com/advanced-go/core/http2/http2test"
	"github.com/advanced-go/core/io2"
	"github.com/advanced-go/core/runtime"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

const (
	emptyEntry = "file://[cwd]/slotest/resource/empty.json"
	validEntry = "file://[cwd]/slotest/resource/slo-entry-v1.json"
)

func Test_httpHandler(t *testing.T) {
	type args struct {
		req    string
		resp   string
		result string
	}
	tests := []struct {
		name string
		args args
	}{
		{"put-entries", args{req: "put-req-v1.txt", resp: "put-resp-v1.txt", result: io2.StatusOK}},
		{"get-entries", args{req: "get-req-v1.txt", resp: "get-resp-v1.txt", result: validEntry}},
		{"delete-entries", args{req: "delete-req-v1.txt", resp: "delete-resp-v1.txt", result: io2.StatusOK}},
	}
	for _, tt := range tests {
		failures, req, resp := http2test.ReadHttp("file://[cwd]/slotest/resource/", tt.args.req, tt.args.resp)
		if failures != nil {
			t.Errorf("ReadHttp() failures = %v", failures)
			continue
		}
		var err error
		req, err = http2test.UpdateUrl(tt.args.result, req)
		if err != nil {
			t.Errorf("UpdateUrl() failure = %v", err)
			continue
		}
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			// ignoring returned status as any errors will be reflected in the response StatusCode
			httpHandler[runtime.Output](w, req)

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
					// test types
					if content {
						if !reflect.DeepEqual(gotT, wantT) {
							t.Errorf("DeepEqual() got = %v, want %v", gotT, wantT)
						}
					}
				}
			}
		})
	}
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

//t.Run(tt.name, func(t *testing.T) {
//	if got := entryHandler(tt.args.w, tt.args.r); !reflect.DeepEqual(got, tt.want) {
//		t.Errorf("entryHandler() = %v, want %v", got, tt.want)
//	}
//})
