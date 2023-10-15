package action

import (
	"context"
	"encoding/json"
	"github.com/go-ai-agent/core/runtime"
	"net/url"
	"time"
)

func GetAction[E runtime.ErrorHandler](ctx context.Context, url *url.URL) ([]byte, *runtime.Status) {
	return nil, runtime.NewStatusOK()
}

type Entry struct {
	CreatedTS  time.Time
	ActivityID string // Some form of UUID
	Agent      string
	Assignment string
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
