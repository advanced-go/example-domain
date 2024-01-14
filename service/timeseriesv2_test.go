package service

import (
	"github.com/advanced-go/core/http2/http2test"
	"github.com/advanced-go/core/runtime"
	"github.com/advanced-go/example-domain/timeseries/entryv2"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

const (
	//emptyEntry = "file://[cwd]/entryv2test/resource/empty.json"
	validV2Entry = "file://[cwd]/servicetest/timeseriesv2/timeseries-entry-v2.json"
)

func Test_timeseriesV2Handler(t *testing.T) {
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
		{"get-entries", args{req: "get-req-v2.txt", resp: "get-resp-v2.txt", result: map[string]string{"getEntries": validV2Entry}}},
		//	{"get-entries-by-controller", args{req: "get-ctrl-req.txt", resp: "get-ctrl-resp.txt",resultemptyEntry}},
		{"delete-entries", args{req: "delete-req-v2.txt", resp: "delete-resp-v2.txt", result: map[string]string{"deleteEntries": ""}}},
	}
	for _, tt := range tests {
		failures, req, resp := http2test.ReadHttp("file://[cwd]/servicetest/timeseriesv2/", tt.args.req, tt.args.resp)
		if failures != nil {
			t.Errorf("ReadHttp() failures = %v", failures)
			continue
		}
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			// ignoring returned status as any errors will be reflected in the response StatusCode
			timeseriesHandlerV2[runtime.Output](w, req)

			// kludge for BUG in response recorder
			w.Result().Header = w.Header()

			// test status code
			if w.Result().StatusCode != resp.StatusCode {
				t.Errorf("StatusCode got = %v, want %v", w.Result().StatusCode, resp.StatusCode)
			} else {
				// test headers if needed - test2.Headers(w.Result(),resp,names... string) (failures []Args)

				// test content size and unmarshal types
				var gotT, wantT []entryv2.Entry
				var content bool
				failures, content, gotT, wantT = http2test.Content[[]entryv2.Entry](w.Result(), resp, timeseriesV2TestBytes)
				if failures != nil {
					timeseriesV2Errorf(t, failures)
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

func timeseriesV2TestBytes(got *http.Response, gotBytes []byte, want *http.Response, wantBytes []byte) []http2test.Args {
	//fmt.Printf("got = %v\n[len:%v]\n", string(gotBytes), len(gotBytes))
	//fmt.Printf("want = %v\n[len:%v]\n", string(wantBytes), len(wantBytes))
	return nil
}

func timeseriesV2Errorf(t *testing.T, failures []http2test.Args) {
	for _, arg := range failures {
		t.Errorf("%v got = %v want = %v", arg.Item, arg.Got, arg.Want)
	}
}
