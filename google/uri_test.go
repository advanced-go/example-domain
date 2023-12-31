package google

import (
	"fmt"
	"net/url"
)

func Example_Resolve_Default() {
	id := searchTag
	uri := resolver.Build(id, nil)
	fmt.Printf("test: Resolve(\"%v\") -> %v\n", id, uri)

	id = searchTag
	v := make(url.Values)
	v.Add("q", "golang")
	uri = resolver.Build(id, v)
	fmt.Printf("test: Resolve(\"%v\") -> %v\n", id, uri)

	//Output:
	//test: Resolve("search") -> https://www.google.com/search
	//test: Resolve("search") -> https://www.google.com/search?q=golang

}

func Example_Override_Host() {
	id := searchTag
	v := make(url.Values)
	v.Add("q", "golang")
	resolver.SetOverride(nil, "http://localhost:8080")

	uri := resolver.Build(id, nil)
	fmt.Printf("test: Resolve(\"%v\") -> %v\n", id, uri)

	id = searchTag
	uri = resolver.Build(id, v)
	fmt.Printf("test: Resolve(\"%v\") -> %v\n", id, uri)

	//Output:
	//test: Resolve("search") -> http://localhost:8080/search
	//test: Resolve("search") -> http://localhost:8080/search?q=golang

}

func testOverrideURL(id string) (string, string) {
	switch id {
	case searchTag:
		return "file://[cwd]/resource/query-result.txt", ""
	}
	return "", ""
}

func Example_Override_URL() {
	id := searchTag
	v := make(url.Values)
	v.Add("q", "golang")
	resolver.SetOverride(testOverrideURL, "")

	uri := resolver.Build(id, nil)
	fmt.Printf("test: Resolve(\"%v\") -> %v\n", id, uri)

	id = searchTag
	uri = resolver.Build(id, v)
	fmt.Printf("test: Resolve(\"%v\") -> %v\n", id, uri)

	//Output:
	//test: Resolve("search") -> file://[cwd]/resource/query-result.txt
	//test: Resolve("search") -> file://[cwd]/resource/query-result.txt

}
