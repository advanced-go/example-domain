package service

import (
	uri2 "github.com/advanced-go/core/uri"
)

const (
	searchTemplate = "/search?%v"
)

var (
	resolver = uri2.NewResolver()
)

func init() {
	resolver.SetTemplates([]uri2.Pair{{searchTemplate, "http://localhost:8081/github/advanced-go/search/provider:search?%v"}})
}
