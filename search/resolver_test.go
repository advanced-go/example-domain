package search

import (
	"fmt"
)

func Example_Build() {
	uri := resolver.Build(authority, searchPath, "")
	fmt.Printf("test: resolver.Build() -> [uri:%v]\n", uri)

	//Output:
	//test: resolver.Build() -> [uri:http://localhost:8081github.com/advanced-go/search/provider:search?]

}
