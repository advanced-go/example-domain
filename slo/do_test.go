package slo

import "fmt"

func Example_PkgUri() {
	fmt.Printf("test: PkgUri = %v\n", PkgUri)
	fmt.Printf("test: Pattern = %v\n", Pattern)

	//Output:
	//test: PkgUri = github.com/go-ai-agent/example-domain/slo
	//test: Pattern = /go-ai-agent/example-domain/slo/

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
