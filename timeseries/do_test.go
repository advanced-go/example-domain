package timeseries

import (
	"fmt"
	"github.com/go-ai-agent/core/log2"
	"github.com/go-ai-agent/core/runtime"
)

func init() {
	BypassLogging()
}

func BypassLogging() {
	wrapper = log2.WrapBypass(newDoHandler[runtime.LogError]())
}

func Example_PkgUri() {
	fmt.Printf("test: PkgUri = %v\n", PkgUri)
	fmt.Printf("test: Pattern = %v\n", Pattern)
	fmt.Printf("test: EntryV1Variant = %v\n", EntryV1Variant)

	//Output:
	//test: PkgUri = github.com/go-ai-agent/example-domain/timeseries
	//test: Pattern = /go-ai-agent/example-domain/timeseries/

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
