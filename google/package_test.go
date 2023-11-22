package google

import (
	"fmt"
	"reflect"
)

func Example_PkgUri() {
	pkgUri2 := reflect.TypeOf(any(pkg{})).PkgPath()
	fmt.Printf("test: PkgPath         = \"%v\"\n", pkgUri2)

	//Output:
	//test: PkgUri         = "github.com/advanced-go/example-domain/google"
	//test: PkgPath        = "/advanced-go/example-domain/google"
	//test: Pattern        = "/advanced-go/example-domain/google/"

}
