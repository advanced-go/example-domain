package activity

import (
	"github.com/go-ai-agent/core/http2"
	"github.com/go-ai-agent/core/log2"
	"github.com/go-ai-agent/core/runtime"
	"net/http"
)

type pkg struct{}

const (
	PkgUri         = "github.com/go-ai-agent/example-domain/activity"
	PkgPath        = "/go-ai-agent/example-domain/activity"
	Pattern        = "/go-ai-agent/example-domain/activity/"
	EntryV1Variant = "github.com/go-ai-agent/example-domain/activity/EntryV1"

	postEntryLoc = PkgUri + "/PostEntry"
)

// GetEntryConstraints - Get constraints
type GetEntryConstraints interface {
	[]EntryV1
}

// GetEntry - generic get function with context and uri for resource selection and filtering
func GetEntry[T GetEntryConstraints](ctx any, uri string) (t T, status *runtime.Status) {
	defer log2.Log(ctx, "GET", uri, log2.NewStatusCodeClosure(&status))()
	return getEntryHandler[T, runtime.LogError](ctx, uri)
}

// PostEntryConstraints - Post constraints
type PostEntryConstraints interface {
	[]EntryV1 | []byte | runtime.Nillable
}

// PostEntry - exchange function
func PostEntry[T PostEntryConstraints](ctx any, method, uri, variant string, body T) (any, *runtime.Status) {
	req, status := http2.NewRequest(ctx, method, uri, variant, nil)
	if !status.OK() {
		var e runtime.LogError
		e.Handle(status, runtime.RequestId(ctx), postEntryLoc)
		return nil, status
	}
	return postWrapper(ctx, req, body)
}

// HttpHandler - Http endpoint
func HttpHandler(w http.ResponseWriter, r *http.Request) {
	httpWrapper(nil, w, r)
}

type EntryV1 struct {
	//CreatedTS    time.Time
	ActivityID   string // Some form of UUID
	ActivityType string // trace|action
	Agent        string
	AgentUri     string // {host}:{agent}

	Assignment  string
	Controller  string
	Behavior    string
	Description string
}
