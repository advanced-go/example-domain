package service

import (
	"fmt"
	"reflect"
)

func Example_PkgUri() {
	pkgUri := reflect.TypeOf(any(pkg{})).PkgPath()
	fmt.Printf("test: PkgPath  = \"%v\"\n", pkgUri)

	//Output:
	//test: PkgPath  = "github.com/advanced-go/example-domain/service"

}
