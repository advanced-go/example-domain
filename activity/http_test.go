package activity

import (
	"fmt"
	"github.com/advanced-go/core/access"
	"github.com/advanced-go/core/http2/http2test"
	"github.com/advanced-go/core/io2"
	"github.com/advanced-go/core/runtime"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func _Example_HttpHandler() {
	access.EnableTestLogHandler()

	addEntry(nil, []Entry{{ActivityID: "activity-uuid",
		ActivityType: "trace",
		Agent:        "agent-controller",
		AgentUri:     "https://host/agent-path",
		Assignment:   "usa:west::test-service:0123456789",
		Controller:   "host-controller",
		Behavior:     "RateLimiting",
		Description:  "Analyzing observation",
	}},
	)

	rec := http2test.NewRecorder()
	req, _ := http.NewRequest("", "https://localhost:8080/advanced-go/example-domain/timeseries/entry", nil)
	//req.Header.Add(http2.ContentLocation, EntryV1Variant)
	HttpHandler(rec, req)
	resp := rec.Result()
	buf, status := io2.ReadAll(resp.Body)
	fmt.Printf("test: HttpHandler() -> [code:%v] [status:%v] [data:%v]\n", rec.Code, status, string(buf))

	//Output:
	//test: HttpHandler() -> 404

}

func resolveUrl(req *http.Request) (*http.Request, error) {
	var err error
	var newUrl string

	switch req.Method {
	case http.MethodGet:
		if req.URL.Query() != nil && len(req.URL.Query().Get(Type)) > 0 {
			newUrl = "file://[cwd]/activitytest/resource/activity-type-entry-v1.json"
		} else {
			newUrl = "file://[cwd]/activitytest/resource/activity-entry-v1.json"
		}
	case http.MethodDelete:
	case http.MethodPut:
		newUrl = "file://[cwd]/activitytest/resource/empty.json"
	}
	//if len(newUrl) == 0 {
	//	return req, nil
	//}
	req.URL, err = url.Parse(newUrl)
	return req, err
}

func Test_httpHandler(t *testing.T) {
	basePath := "file://[cwd]/activitytest/resource/"
	deleteEntries(nil)
	fmt.Printf("test: Start Entries -> %v\n", len(list))
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
		{"get-entries-by-type", args{req: "get-type-req-v1.txt", resp: "get-type-resp-v1.txt"}},
		//{"delete-entries", args{req: "delete-req-v1.txt", resp: "delete-resp-v1.txt"}},
	}
	for _, tt := range tests {
		failures, req, resp := http2test.ReadHttp(basePath, tt.args.req, tt.args.resp)
		if failures != nil {
			t.Errorf("ReadHttp() failures = %v", failures)
			continue
		}
		var err error
		req, err = resolveUrl(req)
		if err != nil {
			t.Errorf("resolveUrl() failure = %v", err)
			continue
		}
		t.Run(tt.name, func(t *testing.T) {
			w := http2test.NewRecorder()
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
