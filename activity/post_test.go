package activity

import (
	"fmt"
	"github.com/advanced-go/core/access"
	"github.com/advanced-go/core/runtime"
	"net/http"
)

func Example_postEntryHandler() {
	req, _ := http.NewRequest("put", "https://www.google.com", nil)
	//if !status.OK() {
	//	fmt.Printf("test: NewRequest() -> [status%v]\n", status)
	//}

	req.Header.Set(runtime.XRequestId, "1234-5678")
	_, status := postEntryHandler(nil, req, nil)
	fmt.Printf("test: postEntryHandler() -> [status:%v]\n", status)

	req, _ = http.NewRequest("PUT", "https://www.google.com", nil)
	//req.Header.Set(ContentLocation, EntryV1Variant)
	req.Header.Set(runtime.XRequestId, "8765-4321")
	_, status = postEntryHandler(nil, req, "invalid string type")
	fmt.Printf("test: postEntryHandler() -> [status:%v]\n", status)

	//Output:
	//test: postEntryHandler() -> [status:Invalid Argument [error invalid variant: [<empty>] for [github.com/advanced-go/example-domain/activity]]]
	//test: postEntryHandler() -> [status:Invalid Content [invalid body type: string]]

}

func Example_PostEntry() {
	access.EnableDebugLogHandler()
	entries := []Entry{
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
	_, status := PostEntry[[]Entry](h, "PUT", "http://localhost:8080/advanced-go/example-domain/activity", entries)
	fmt.Printf("test: PostEntry() -> [status:%v]\n", status)

	//Output:
	//{ "activity": "trace" "agent": "agent-test"  "controller": "controller-test"  "message": "desc-1"  }
	//{ "activity": "trace" "agent": "agent-test"  "controller": "controller-test"  "message": "desc-2"  }
	//function StatusOK.AddLocation() is not implemented
	//{ "traffic":"internal", "start":2023-11-20 11:34:15.729060, "duration":0, "request-id":"123-456", "protocol":"HTTP/1.1", "method":"PUT", "url":"http://localhost:8080/advanced-go/example-domain/activity", "host":"localhost:8080", "path":"/advanced-go/example-domain/activity", "status-code":200, "threshold":-1, "threshold-flags":null }
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
