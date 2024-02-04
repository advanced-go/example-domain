package service

import (
	"github.com/advanced-go/core/http2/http2test"
	"github.com/advanced-go/core/io2"
	"github.com/advanced-go/core/runtime"
	"github.com/advanced-go/example-domain/activity"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

const (
	sloValidEntry = "file://[cwd]/servicetest/slo/slo-entry-v1.json"
)

func Test_sloHandler(t *testing.T) {
	basePath := "file://[cwd]/servicetest/slo/"
	type args struct {
		req    string
		resp   string
		result any
	}
	tests := []struct {
		name string
		args args
	}{
		{"put-entries", args{req: "put-req-v1.txt", resp: "put-resp-v1.txt", result: map[string]string{"addEntries": io2.StatusOKUri}}},
		{"get-entries", args{req: "get-req-v1.txt", resp: "get-resp-v1.txt", result: map[string]string{"getEntries": sloValidEntry}}},
		{"delete-entries", args{req: "delete-req-v1.txt", resp: "delete-resp-v1.txt", result: map[string]string{"deleteEntries": ""}}},
	}
	for _, tt := range tests {
		failures, req, resp := http2test.ReadHttp(basePath, tt.args.req, tt.args.resp)
		if failures != nil {
			t.Errorf("ReadHttp() failures = %v", failures)
			continue
		}
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			// ignoring returned status as any errors will be reflected in the response StatusCode
			sloHandler[runtime.Output](w, req)

			// test status code
			if w.Result().StatusCode != resp.StatusCode {
				t.Errorf("StatusCode got = %v, want %v", w.Result().StatusCode, resp.StatusCode)
			} else {
				// test headers if needed - test2.Headers(w.Result(),resp,names... string) (failures []Args)

				// test content size and unmarshal types
				var gotT, wantT []activity.EntryV1
				var content bool
				failures, content, gotT, wantT = http2test.Content[[]activity.EntryV1](w.Result(), resp, sloTestBytes)
				if failures != nil {
					sloErrorf(t, failures)
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
}

func sloTestBytes(got *http.Response, gotBytes []byte, want *http.Response, wantBytes []byte) []http2test.Args {
	//fmt.Printf("got = %v\n[len:%v]\n", string(gotBytes), len(gotBytes))
	//fmt.Printf("want = %v\n[len:%v]\n", string(wantBytes), len(wantBytes))
	return nil
}

func sloErrorf(t *testing.T, failures []http2test.Args) {
	for _, arg := range failures {
		t.Errorf("%v got = %v want = %v", arg.Item, arg.Got, arg.Want)
	}
}

//t.Run(tt.name, func(t *testing.T) {
//	if got := entryHandler(tt.args.w, tt.args.r); !reflect.DeepEqual(got, tt.want) {
//		t.Errorf("entryHandler() = %v, want %v", got, tt.want)
//	}
//})
