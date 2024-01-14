package service

import (
	"github.com/advanced-go/core/http2/http2test"
	"github.com/advanced-go/core/runtime"
	"github.com/advanced-go/example-domain/activity"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

const (
	activityStateEntry     = "file://[cwd]/servicetest/activity/activity-entry-v1.json"
	activityStateEntryType = "file://[cwd]/servicetest/activity/activity-type-entry-v1.json"
	activityStateEmpty     = "file://[cwd]/servicetest/activity/empty.json"
)

func Test_activityHandler(t *testing.T) {
	basePath := "file://[cwd]/servicetest/activity/"
	type args struct {
		req    string
		resp   string
		result any
	}
	tests := []struct {
		name string
		args args
	}{
		{"get-entries-empty", args{req: "get-req-v1.txt", resp: "get-resp-v1-empty.txt", result: map[string]string{"getEntries": activityStateEmpty}}},
		{"put-entries", args{req: "put-req-v1.txt", resp: "put-resp-v1.txt", result: map[string]string{"addEntries": ""}}},
		{"get-entries", args{req: "get-req-v1.txt", resp: "get-resp-v1.txt", result: map[string]string{"getEntries": activityStateEntry}}},
		{"get-entries-by-type", args{req: "get-type-req-v1.txt", resp: "get-type-resp-v1.txt", result: map[string]string{"getEntriesByType": activityStateEntryType}}},
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
			activityHandler[runtime.Output](w, req)

			// kludge for BUG in response recorder
			w.Result().Header = w.Header()

			// test status code
			if w.Result().StatusCode != resp.StatusCode {
				t.Errorf("StatusCode got = %v, want %v", w.Result().StatusCode, resp.StatusCode)
			} else {
				// test headers if needed - test2.Headers(w.Result(),resp,names... string) (failures []Args)

				// test content size and unmarshal types
				var gotT, wantT []activity.EntryV1
				var content bool
				failures, content, gotT, wantT = http2test.Content[[]activity.EntryV1](w.Result(), resp, activityTestBytes)
				if failures != nil {
					activityErrorf(t, failures)
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

func activityTestBytes(got *http.Response, gotBytes []byte, want *http.Response, wantBytes []byte) []http2test.Args {
	//fmt.Printf("got = %v\n[len:%v]\n", string(gotBytes), len(gotBytes))
	//fmt.Printf("want = %v\n[len:%v]\n", string(wantBytes), len(wantBytes))
	return nil
}

func activityErrorf(t *testing.T, failures []http2test.Args) {
	for _, arg := range failures {
		t.Errorf("%v got = %v want = %v", arg.Item, arg.Got, arg.Want)
	}
}

//t.Run(tt.name, func(t *testing.T) {
//	if got := entryHandler(tt.args.w, tt.args.r); !reflect.DeepEqual(got, tt.want) {
//		t.Errorf("entryHandler() = %v, want %v", got, tt.want)
//	}
//})
