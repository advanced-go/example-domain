package google

import (
	"github.com/go-ai-agent/core/runtime"
	"net/http"
	"reflect"
	"testing"
)

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

func Test_searchHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
		want *runtime.Status
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchHandler[runtime.DebugError](tt.args.w, tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
