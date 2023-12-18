package timeseriesvar

import (
	"fmt"
	"github.com/advanced-go/core/runtime"
	"net/http"
	"net/url"
	"reflect"
	"time"
)

var entries = []EntryV1{{
	Traffic:   "ingress",
	Start:     time.Now().UTC(),
	Duration:  750,
	RequestId: "98765",
	Url:       "https://www.somestupiddomain.com/help",
}}

func Example_PkgUri() {
	pkgUri2 := reflect.TypeOf(any(pkg{})).PkgPath()
	//pkgPath2 := runtime.PathFromUri(pkgUri2)
	entryV1 := pkgUri2 + "/" + reflect.TypeOf(EntryV1{}).Name()
	entryV2 := pkgUri2 + "/" + reflect.TypeOf(EntryV2{}).Name()

	fmt.Printf("test: PkgPath        = \"%v\"\n", pkgUri2)
	fmt.Printf("test: EntryV1Variant = \"%v\"\n", entryV1)
	fmt.Printf("test: EntryV2Variant = \"%v\"\n", entryV2)

	//Output:
	//test: PkgPath        = "github.com/advanced-go/example-domain/timeseriesvar"
	//test: EntryV1Variant = "github.com/advanced-go/example-domain/timeseriesvar/EntryV1"
	//test: EntryV2Variant = "github.com/advanced-go/example-domain/timeseriesvar/EntryV2"

}

func getProxy(h http.Header, uri *url.URL) (any, runtime.Status) {
	fmt.Printf("test: getProxy() -> in proxy\n")
	return entries, runtime.NewStatusOK() //http.StatusInternalServerError)
}

func _Example_GetWithProxy() {
	//ctx := runtime.NewRequestIdContext(nil, "get-654-321")
	//ctx = runtime.NewProxyContext(ctx, getProxy)
	h := make(http.Header)
	u, _ := url.Parse("https://google.com/search")
	e, status := getEntryHandler[[]EntryV1](nil, h, u)
	fmt.Printf("test: getEntryHandler[[]EntryV1]() -> [status:%v] [entries:%v]\n", status, len(e))

	//Output:
	//test: getProxy() -> in proxy
	//test: getEntryHandler[[]EntryV1]() -> [status:OK] [entries:1]

}
