# example-domain

Provides common functionality utilized by an AI agent managing service resiliency. The packages, activity, slo, and timeseries, provide uniform interfaces in their respective package.go.

~~~
// Resource identifier
const (
    PkgPath = "github/advanced-go/example-domain/activity"
)

// Get - get entries
func Get(h http.Header, uri string) (entries []Entry, status runtime.Status) {
 // implementation details
}

// PostConstraints - Post constraints
type PostConstraints interface {
	[]Entry | []byte | runtime.Nillable
}

// Post - exchange function for POST, PUT, DELETE...
func Post[T PostConstraints](h http.Header, method, uri string, body T) (t any, status runtime.Status) {
 // implementation details
}
~~~

The service package service is the resource that handles incoming HTTP requests. There are seperate handlers for each supporting package plus a search endpoint. Testing is implemented via the uniform HttpHandler interface in each handler.

The following uniform interfaces are provided:

~~~
// Resource identifier
const (
    PkgPath = "github/advanced-go/example-domain/service"
)

// HttpHandler - http endpoint
func HttpHandler(w http.ResponseWriter, r *http.Request) {
 // implementation details
}
~~~
   


Access logging is implemented by a call to a uniform logging interface in core/access
~~~
defer access.LogDeferred(access.InternalTraffic, access.NewRequest(h, http.MethodGet, getLoc), -1, "", access.NewStatusCodeClosure(&status))()
return getHandler[runtime.LogError](nil, h, u)
~~~

## activity
[Activity][activitypkg] provides an audit trail for all AI agent actions, interactions, and results of analysis. 

## slo
[SLO][slopkg] implements SLO's. 

## timeseries
[Timeseries][timeseriespkg] implements versioned access log events. 

[activitypkg]: <https://pkg.go.dev/github.com/advanced-go/example-domain/activity>
[slopkg]: <https://pkg.go.dev/github.com/advanced-go/example-domain/slo>
[timeseriespkg]: <https://pkg.go.dev/github.com/advanced-go/example-domain/timeseries>
[rfc2626]: <https://datatracker.ietf.org/doc/html/rfc2616>



