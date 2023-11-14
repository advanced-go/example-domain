package activity

import (
	"fmt"
	"github.com/go-ai-agent/core/runtime"
	"reflect"
)

func Example_PkgUri() {
	pkgUri2 := reflect.TypeOf(any(pkg{})).PkgPath()
	pkgPath2 := runtime.PathFromUri(pkgUri2)
	entryV1 := pkgUri2 + "/" + reflect.TypeOf(EntryV1{}).Name()

	fmt.Printf("test: PkgUri         = \"%v\"\n", pkgUri2)
	fmt.Printf("test: PkgPath        = \"%v\"\n", pkgPath2)
	fmt.Printf("test: Pattern        = \"%v\"\n", pkgPath2+"/")
	fmt.Printf("test: EntryV1Variant = \"%v\"\n", entryV1)

	//Output:
	//test: PkgUri         = "github.com/go-ai-agent/example-domain/activity"
	//test: PkgPath        = "/go-ai-agent/example-domain/activity"
	//test: Pattern        = "/go-ai-agent/example-domain/activity/"
	//test: EntryV1Variant = "github.com/go-ai-agent/example-domain/activity/EntryV1"

}
