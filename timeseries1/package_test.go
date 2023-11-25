package timeseries1

import (
	"fmt"
	"reflect"
	"time"
)

var entries = []Entry{{
	Traffic:   "ingress",
	Start:     time.Now().UTC(),
	Duration:  750,
	RequestId: "98765",
	Url:       "https://www.somestupiddomain.com/help",
}}

func Example_PkgUri() {
	pkgUri2 := reflect.TypeOf(any(pkg{})).PkgPath()

	fmt.Printf("test: PkgPath = \"%v\"\n", pkgUri2)

	//Output:
	//test: PkgPath = "github.com/advanced-go/example-domain/timeseries1"

}
