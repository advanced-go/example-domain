package service

import (
	"fmt"
	"github.com/advanced-go/core/runtime"
	uri2 "github.com/advanced-go/core/uri"
	"net/url"
)

const (
	queryArg = "q"
)

func ExampleBuild() {
	v := make(url.Values)
	v.Add(queryArg, "golang")

	uri := resolver.Build(searchTemplate, v.Encode())
	fmt.Printf("test: resolver.Build-Debug(\"%v\") -> [uri:%v]\n", searchTemplate, uri)

	//Output:
	//test: resolver.Build-Debug("/search?%v") -> [uri:http://localhost:8081/github/advanced-go/search/provider:search?q=golang]

}

func ExampleBuild_Override() {
	runtime.SetProdEnvironment()

	v := make(url.Values)
	v.Add(queryArg, "golang")

	uri := resolver.Build(searchTemplate, v.Encode())
	fmt.Printf("test: resolver.Build(\"%v\") -> [uri:%v]\n", searchTemplate, uri)

	resolver.SetTemplates([]uri2.Pair{{searchTemplate, "https://www.google.com/search?q=Pascal"}})
	s := v.Encode()
	uri = resolver.Build(searchTemplate, s)
	fmt.Printf("test: resolver.Build(\"%v\") -> [uri:%v]\n", searchTemplate, uri)

	resolver.SetTemplates([]uri2.Pair{{searchTemplate, "file://[cwd]/providertest/resource/query-result.txt"}})
	s = v.Encode()
	uri = resolver.Build(searchTemplate, s)
	fmt.Printf("test: resolver.Build(\"%v\") -> [uri:%v]\n", searchTemplate, uri)

	//Output:
	//test: resolver.Build("/search?%v") -> [uri:http://localhost:8081/github/advanced-go/search/provider:search?q=golang]
	//test: resolver.Build("/search?%v") -> [uri:https://www.google.com/search?q=Pascal]
	//test: resolver.Build("/search?%v") -> [uri:file://[cwd]/providertest/resource/query-result.txt]

}
