package service

import (
	uri2 "github.com/advanced-go/stdlib/uri"
)

const (
	searchTemplate = "/search?%v"
)

var (
	resolver = uri2.NewResolver()
)

func init() {
	resolver.SetTemplates([]uri2.Attr{{searchTemplate, "http://localhost:8081/github/advanced-go/search/provider:search?%v"}})
}
