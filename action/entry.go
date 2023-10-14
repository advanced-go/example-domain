package action

import (
	"context"
	"encoding/json"
	"github.com/go-ai-agent/core/runtime"
	"net/url"
)

func GetAction[E runtime.ErrorHandler](ctx context.Context, url *url.URL) ([]byte, *runtime.Status) {
	return nil, runtime.NewStatusOK()
}

type Entry struct {
	Controller string
	Behavior   string
	Action     string
}

var list []Entry

func GetEntries() ([]byte, error) {
	return json.Marshal(list)
}

func AddEntry(e Entry) {
	list = append(list, e)
}
