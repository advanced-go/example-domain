package activity

import (
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"net/http"
)

func Example_postEntryHandler() {
	req, _ := http.NewRequest(http.MethodPut, "https://www.google.com", nil)
	//if !status.OK() {
	//	fmt.Printf("test: NewRequest() -> [status%v]\n", status)
	//}

	req.Header.Set(httpx.XRequestId, "1234-5678")
	_, status := postEntryHandler[core.Output](nil, req.Header, req.Method, nil, nil)
	fmt.Printf("test: postEntryHandler() -> [status:%v]\n", status)

	req, _ = http.NewRequest("PUT", "https://www.google.com", nil)
	//req.Header.Set(ContentLocation, EntryV1Variant)
	req.Header.Set(httpx.XRequestId, "8765-4321")
	_, status = postEntryHandler[core.Output](nil, req.Header, req.Method, nil, "invalid string type")
	fmt.Printf("test: postEntryHandler() -> [status:%v]\n", status)

	//Output:
	//test: postEntryHandler() -> [status:Invalid Content]
	//{ "code":90, "status":"Invalid Content", "request-id":"8765-4321", "errors" : [ "invalid body type: string" ], "trace" : [ "https://github.com/advanced-go/example-domain/tree/main/activity#postEntryHandler[...]","https://github.com/advanced-go/example-domain/tree/main/activity#createEntries" ] }
	//test: postEntryHandler() -> [status:Invalid Content [invalid body type: string]]

}

/*

func Example_PostEntry() {
	access.EnableTestLogger()
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
	h.Add(core.XRequestId, "123-456")
	_, status := PostEntry[[]Entry](h, "PUT", nil, entries)
	//_, status := PostEntry[[]Entry](h, "PUT", "http://localhost:8080/advanced-go/example-domain/activity", entries)

	fmt.Printf("test: PostEntry() -> [status:%v]\n", status)

	//Output:
	//{ "activity": "trace" "agent": "agent-test"  "controller": "controller-test"  "message": "desc-1"  }
	//{ "activity": "trace" "agent": "agent-test"  "controller": "controller-test"  "message": "desc-2"  }
	//test: PostEntry() -> [status:OK]

}

*/

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
