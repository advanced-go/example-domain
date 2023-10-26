package google

import "fmt"

func Example_PkgUri() {
	fmt.Printf("test: PkgUrl %v\n", PkgUrl)
	fmt.Printf("test: PkgUri %v\n", PkgUri)
	fmt.Printf("test: SearchPath %v\n", SearchPath)

	//Output:
	//test: PkgUrl file://github.com/go-ai-agent/example-domain/slo
	//test: PkgUri github.com/go-ai-agent/example-domain/slo
	//test: EntryPath /go-ai-agent/example-domain/slo/entry

}
