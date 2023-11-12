package timeseries

import (
	"fmt"
	"net/http"
)

func Example_PkgUri() {
	fmt.Printf("test: PkgUri = %v\n", PkgUri)
	fmt.Printf("test: Pattern = %v\n", Pattern)
	fmt.Printf("test: EntryV1Variant = %v\n", EntryV1Variant)

	//Output:
	//test: PkgUri = github.com/go-ai-agent/example-domain/timeseries
	//test: Pattern = /go-ai-agent/example-domain/timeseries/
	//test: EntryV1Variant = github.com/go-ai-agent/example-domain/timeseries/EntryV1

}

func Example_verifyVariant() {
	uri := "https://www/google/com"
	r, _ := http.NewRequest("", uri, nil)
	v := verifyVariant(r)
	fmt.Printf("test: verifyVariant(%v) -> %v\n", uri, v)

	uri = "https://www/google/com?q=golang"
	r, _ = http.NewRequest("", uri, nil)
	v = verifyVariant(r)
	fmt.Printf("test: verifyVariant(%v) -> %v\n", uri, v)

	uri = "https://www/google/com?v=3"
	r, _ = http.NewRequest("", uri, nil)
	v = verifyVariant(r)
	fmt.Printf("test: verifyVariant(%v) -> %v\n", uri, v)

	uri = "https://www/google/com?v=1"
	r, _ = http.NewRequest("", uri, nil)
	v = verifyVariant(r)
	fmt.Printf("test: verifyVariant(%v) -> %v\n", uri, v)

	uri = "https://www/google/com?v=2"
	r, _ = http.NewRequest("", uri, nil)
	v = verifyVariant(r)
	fmt.Printf("test: verifyVariant(%v) -> %v\n", uri, v)

	//Output:
	//test: verifyVariant(https://www/google/com) -> github.com/go-ai-agent/example-domain/timeseries/EntryV1
	//test: verifyVariant(https://www/google/com?q=golang) -> github.com/go-ai-agent/example-domain/timeseries/EntryV1
	//test: verifyVariant(https://www/google/com?v=3) -> github.com/go-ai-agent/example-domain/timeseries/EntryV1
	//test: verifyVariant(https://www/google/com?v=1) -> github.com/go-ai-agent/example-domain/timeseries/EntryV1
	//test: verifyVariant(https://www/google/com?v=2) -> github.com/go-ai-agent/example-domain/timeseries/EntryV2

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
