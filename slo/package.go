package slo

import (
	"github.com/go-ai-agent/core/http2"
	"github.com/go-ai-agent/core/log2"
	"github.com/go-ai-agent/core/runtime"
	"net/http"
	"time"
)

type pkg struct{}

const (
	PkgUri         = "github.com/go-ai-agent/example-domain/slo"
	PkgPath        = "/go-ai-agent/example-domain/slo"
	Pattern        = "/go-ai-agent/example-domain/slo/"
	EntryV1Variant = "github.com/go-ai-agent/example-domain/slo/EntryV1"

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
	var e runtime.LogError

	req, status := http2.NewRequest(ctx, method, uri, variant, nil)
	if !status.OK() {
		e.Handle(status, runtime.RequestId(ctx), postEntryLoc)
		return nil, status
	}
	return postWrapper(ctx, req, body)
}

// HttpHandler - http endpoint
func HttpHandler(w http.ResponseWriter, r *http.Request) {
	httpWrapper(nil, w, r)
}

type EntryV1 struct {
	CreatedTS time.Time
	Id        string
	// What does this apply to
	Controller string

	// Types of SLOs
	// percentage of traffic : 10% or 10
	// latency percentile: 99/500ms
	Threshold   string // Either percentage of traffic, or latency percentile
	StatusCodes string // For percentage
}
