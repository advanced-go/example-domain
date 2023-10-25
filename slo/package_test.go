package slo

import (
	"fmt"
)

func Example_PkgUri() {
	fmt.Printf("test: PkgUrl %v\n", PkgUrl)
	fmt.Printf("test: PkgUri %v\n", PkgUri)
	fmt.Printf("test: SLOPath %v\n", SLOPath)

	//Output:
	//test: PkgUrl file://github.com/go-ai-agent/example-domain/slo
	//test: PkgUri github.com/go-ai-agent/example-domain/slo
	//test: SLOPath /go-ai-agent/example-domain/slo/entry

}

/*
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

func Test_sloHandler(t *testing.T) {
	type args struct {
		req  string
		resp string
	}
	tests := []struct {
		name string
		args args
	}{
		{"get entries", args{req: "get-req.txt", resp: "get-resp.txt"}},
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
			sloHandler[runtime.BypassError](w, req)

			// test status code
			if w.Result().StatusCode != resp.StatusCode {
				t.Errorf("StatusCode got = %v, want %v", w.Result().StatusCode, resp.StatusCode)
			}

			// test headers if needed - test2.Headers(w.Result(),resp,names... string) (failures []Args)

			// test content size and unmarshal types
			var gotT, wantT []SLO
			failures, content, gotT, wantT = httpxtest.Content[[]SLO](w.Result(), resp)
			if failures != nil {
				t.Errorf("Content() failures = %v", failures)
			}

			// test types
			if content {
				if !reflect.DeepEqual(gotT, wantT) {
					t.Errorf("DeepEqual() got = %v, want %v", gotT, wantT)
				}
			}
		})
	}
}

//t.Run(tt.name, func(t *testing.T) {
//	if got := entryHandler(tt.args.w, tt.args.r); !reflect.DeepEqual(got, tt.want) {
//		t.Errorf("entryHandler() = %v, want %v", got, tt.want)
//	}
//})


*/
