package google

import (
	"fmt"
	"github.com/advanced-go/core/runtime"
	"reflect"
)

func Example_PkgUri() {
	pkgUri2 := reflect.TypeOf(any(pkg{})).PkgPath()
	pkgPath2 := runtime.PathFromUri(pkgUri2)

	fmt.Printf("test: PkgUri         = \"%v\"\n", pkgUri2)
	fmt.Printf("test: PkgPath        = \"%v\"\n", pkgPath2)
	fmt.Printf("test: Pattern        = \"%v\"\n", pkgPath2+"/")

	//Output:
	//test: PkgUri         = "github.com/advanced-go/example-domain/google"
	//test: PkgPath        = "/advanced-go/example-domain/google"
	//test: Pattern        = "/advanced-go/example-domain/google/"

}
