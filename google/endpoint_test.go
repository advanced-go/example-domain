package google

import (
	"fmt"
	"net/url"
)

func Example_searchEndpoint() {
	endpoint := "/google/search"
	googleUri := "https://www.google.com/search"

	s := PkgUri
	uri, _ := url.Parse(s)

	result := searchUri(uri, endpoint)
	fmt.Printf("test: searchUrl(%v) %v\n", s, result)

	result = searchUri(uri, googleUri)
	fmt.Printf("test: searchUrl(%v) %v\n", s, result)

	s = PkgUri + "?q=testrlz=1C1CHBF"
	uri, _ = url.Parse(s)

	result = searchUri(uri, endpoint)
	fmt.Printf("test: searchUrl(%v) %v\n", s, result)

	result = searchUri(uri, googleUri)
	fmt.Printf("test: searchUrl(%v) %v\n", s, result)

	//Output:
	//test: searchUrl(github.com/advanced-go/example-domain/google) /google/search
	//test: searchUrl(github.com/advanced-go/example-domain/google) https://www.google.com/search
	//test: searchUrl(github.com/advanced-go/example-domain/google?q=testrlz=1C1CHBF) /google/search?q=testrlz=1C1CHBF
	//test: searchUrl(github.com/advanced-go/example-domain/google?q=testrlz=1C1CHBF) https://www.google.com/search?q=testrlz=1C1CHBF

}
