package google

import (
	"fmt"
	"github.com/go-ai-agent/core/exchange"
	"net/url"
)

func Example_PkgUri() {
	fmt.Printf("test: PkgUri = %v\n", pkgUri)
	fmt.Printf("test: SearchEndpoint = %v\n", SearchEndpoint)

	//Output:
	//test: PkgUri = github.com/go-ai-agent/example-domain/google
	//test: SearchEndpoint = /go-ai-agent/example-domain/google/search

}

func Example_searchEndpoint() {
	s := pkgUri
	uri, _ := url.Parse(s)
	result := searchEndpoint(uri)

	fmt.Printf("test: searchEndpoint(%v) %v\n", s, result)

	s = pkgUri + "?q=test&rlz=1C1CHBF"
	uri, _ = url.Parse(s)
	result = searchEndpoint(uri)

	fmt.Printf("test: searchEndpoint(%v) %v\n", s, result)

	//Output:
	//test: searchEndpoint(github.com/go-ai-agent/example-domain/google) /search
	//test: searchEndpoint(github.com/go-ai-agent/example-domain/google?q=test&rlz=1C1CHBF) /search?q=test

}

func Example_Resolve() {
	p := "/go-ai-agent/example-domain/google?q=test&rlz=1C1CHBF"
	uri, _ := url.Parse(p)
	s := exchange.Resolve(searchEndpoint(uri))

	fmt.Printf("test: Resolve(%v) path = %v\n", p, s)

	//Output:
	//test: Resolve(/go-ai-agent/example-domain/google?q=test&rlz=1C1CHBF) path = http://localhost:8080/search?q=test

}
