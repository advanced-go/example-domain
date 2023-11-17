package activity

import (
	"fmt"
	"github.com/advanced-go/core/access"
	"github.com/advanced-go/core/http2"
	"github.com/advanced-go/core/runtime"
	"github.com/advanced-go/core/runtime/runtimetest"
	"net/http"
)

func Example_postEntryHandler() {
	req, status := http2.NewRequest(nil, "put", "", "", nil)
	if !status.OK() {
		fmt.Printf("test: NewRequest() -> [status%v]\n", status)
	}

	req.Header.Set(runtime.XRequestId, "1234-5678")
	_, status = postEntryHandler[runtimetest.DebugError](nil, req, nil)
	fmt.Printf("test: postEntryHandler() -> [status:%v]\n", status)

	req, status = http2.NewRequest(nil, "put", "", EntryV1Variant, nil)
	req.Header.Set(runtime.XRequestId, "8765-4321")
	_, status = postEntryHandler[runtimetest.DebugError](nil, req, "invalid string type")
	fmt.Printf("test: postEntryHandler() -> [status:%v]\n", status)

	//Output:
	//{ "code":3, "status":"Invalid Argument", "id":"1234-5678", "trace" : [ "github.com/advanced-go/example-domain/activity/postEntryHandler","github.com/advanced-go/example-domain/activity/validateVariant" ], "err" : [ "error invalid variant: [<empty>] for [github.com/advanced-go/example-domain/activity]" ] }
	//test: postEntryHandler() -> [status:Invalid Argument [error invalid variant: [<empty>] for [github.com/advanced-go/example-domain/activity]]]
	//{ "code":90, "status":"Invalid Content", "id":"8765-4321", "trace" : [ "github.com/advanced-go/example-domain/activity/postEntryHandler","github.com/advanced-go/example-domain/activity/putEntry" ], "err" : [ "invalid body type: string" ] }
	//test: postEntryHandler() -> [status:Invalid Content [invalid body type: string]]

}

func Example_PostEntry() {
	access.EnableDebugLogHandler()
	entries := []EntryV1{
		{
			ActivityID:   "",
			ActivityType: "trace",
			Agent:        "agent-test",
			AgentUri:     "",
			Assignment:   "",
			Controller:   "controller-test",
			Behavior:     "",
			Description:  "desc-1",
		}, {
			ActivityID:   "",
			ActivityType: "trace",
			Agent:        "agent-test",
			AgentUri:     "",
			Assignment:   "",
			Controller:   "controller-test",
			Behavior:     "",
			Description:  "desc-2",
		}}

	h := make(http.Header)
	h.Add(runtime.XRequestId, "123-456")
	_, status := PostEntry[[]EntryV1](h, "PUT", "http://localhost:8080/advanced-go/example-domain/activity", EntryV1Variant, entries)
	fmt.Printf("test: PostEntry() -> [status:%v]\n", status)

	//Output:
	//{ "activity": "trace" "agent": "agent-test"  "controller": "controller-test"  "message": "desc-1"  }
	//{ "activity": "trace" "agent": "agent-test"  "controller": "controller-test"  "message": "desc-2"  }
	//test: PostEntry() -> [status:OK]

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
