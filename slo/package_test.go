package slo

import (
	"fmt"
	"reflect"
)

func Example_PkgUri() {
	pkgUri2 := reflect.TypeOf(any(pkg{})).PkgPath()
	fmt.Printf("test: PkgPath = \"%v\"\n", pkgUri2)

	//Output:
	//test: PkgPath = "github.com/advanced-go/example-domain/slo"

}
