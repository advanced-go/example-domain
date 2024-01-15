package service

import (
	"github.com/advanced-go/core/runtime"
	uri2 "github.com/advanced-go/core/uri"
)

const (
	searchTemplate = "/search?%v"
)

var (
	resolver = uri2.NewResolver()
)

func init() {
	resolver.SetOverrides([]runtime.Pair{{searchTemplate, "http://localhost:8081/github/advanced-go/search/provider:search?%v"}})
}
