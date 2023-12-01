# example-domain

Provides common functionality utilized by an AI agent managing service resiliency. The packages provide 2 interfaces for integration:
  1. Direct exchange functions - Get and Post, with generic constraints for Post
~~~
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

  3. HTTP handler - implementing http.Handler
~~~
// HttpHandler - http endpoint
func HttpHandler(w http.ResponseWriter, r *http.Request) {
 // implementation details
}
~~~
   
The implementation for the above interfaces, and any additional information needed for integration with the package, are provied in package.go. The timeseries package also provides an implementation for resource versioning. Access logging is supported

Applications that want to use example-domain functionality can integrate directly, by calling the package's functions, or access the functionality hosted in another service, via HTTP. Hosting example-domain packages only requires registering a ServMux handler and pattern, which are both defined in the package.go file. All of the testing, including the Http handler, is automated in the package and does not need to be implemented in a service host. This allows the packages to be deployed in multiple hosts, providing flexibility when creating new functionality. New services can utilize existing services, or integrate directly with the packaged functionality. 

## action
[Action][actionpkg] implements actions that an AI agent can take to affect change in response to an observation. 

## activity
[Activity][activitypkg] provides an audit trail for all AI agent actions, interactions, and results of analysis. 

## google
[Google][googlepkg] provides Google search functionality. 

## slo
[SLO][slopkg] implements SLO's. 

## timeseries
[Timeseries][timeseriespkg] implements versioned access log events. 

## timeseriesvar
[Timeseriesvar][timeseriesvarpkg] implements versioned access log events via [variants][rfc2626]. 



[actionpkg]: <https://pkg.go.dev/github.com/advanced-go/example-domain/action>
[activitypkg]: <https://pkg.go.dev/github.com/advanced-go/example-domain/activity>
[googlepkg]: <https://pkg.go.dev/github.com/advanced-go/example-domain/google>
[slopkg]: <https://pkg.go.dev/github.com/advanced-go/example-domain/slo>
[timeseriespkg]: <https://pkg.go.dev/github.com/advanced-go/example-domain/timeseries>
[timeseriesvarpkg]: <https://pkg.go.dev/github.com/advanced-go/example-domain/timeseriesvar>
[rfc2626]: <https://datatracker.ietf.org/doc/html/rfc2616>



