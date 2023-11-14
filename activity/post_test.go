package activity

import (
	"fmt"
	"github.com/go-ai-agent/core/http2"
	"github.com/go-ai-agent/core/runtime/runtimetest"
)

func Example_PkgUri() {
	fmt.Printf("test: PkgUri = %v\n", PkgUri)
	fmt.Printf("test: Pattern = %v\n", Pattern)
	fmt.Printf("test: EntryV1Variant = %v\n", EntryV1Variant)

	//Output:
	//test: PkgUri = github.com/go-ai-agent/example-domain/activity
	//test: Pattern = /go-ai-agent/example-domain/activity/
	//test: EntryV1Variant = github.com/go-ai-agent/example-domain/activity/EntryV1

}

func Example_postHandler() {
	req, status := http2.NewRequest(nil, "put", "", "")
	if !status.OK() {
		fmt.Printf("test: NewRequest() -> [status%v]\n", status)
	}

	_, status = postHandler[runtimetest.DebugError](nil, req, nil)
	fmt.Printf("test: postHandler() -> %v\n", status)

	req, status = http2.NewRequest(nil, "put", "", EntryV1Variant)
	_, status = postHandler[runtimetest.DebugError](nil, req, "invalid string type")
	fmt.Printf("test: postHandler() -> %v\n", status)

	//Output:
	//{ "code":90, "status":"Invalid Content", "id":"b7d1c98c-808f-11ee-962d-00a55441ed8b", "trace" : [ "","github.com/go-ai-agent/example-domain/activity/doHandler" ], "err" : [ "invalid body type: <nil>" ] }
	//test: postHandler() -> Invalid Content [invalid body type: <nil>]
	//{ "code":90, "status":"Invalid Content", "id":"b7d2a698-808f-11ee-962d-00a55441ed8b", "trace" : [ "","github.com/go-ai-agent/example-domain/activity/doHandler" ], "err" : [ "invalid body type: string" ] }
	//test: postHandler() -> Invalid Content [invalid body type: string]

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

*/
