# example-domain

Provides common functionality utilized by an AI agent managing service resiliency. The packages provide 2 interfaces for integration:
  1. Direct exchange functions - Get and Post, with generic constraints for Post.
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
   
A package.go file implements the above interfaces, and provides any additional type declarations for integration. 

Example-domain packages also address the following functional areas:
1. Resource versioning - the timeseries package supports versioning via seperate sub packages, and is imelemented in package.go
 ~~~
// GetEntryV1 - get entries
func GetEntryV1(h http.Header, uri string) (entries []entryv1.Entry, status runtime.Status) {
	return entryv1.Get(h, uri)
}

// GetEntryV2 - get entries
func GetEntryV2(h http.Header, uri string) (entries []entryv2.Entry, status runtime.Status) {
	return entryv2.Get(h, uri)
}

// PostEntryV1 - exchange function
func PostEntryV1[T entryv1.PostConstraints](h http.Header, method, uri string, body T) (t any, status runtime.Status) {
	return entryv1.Post[T](h, method, uri, body)
}

// PostEntryV2 - exchange function
func PostEntryV2[T entryv2.PostConstraints](h http.Header, method, uri string, body T) (t any, status runtime.Status) {
	return entryv2.Post[T](h, method, uri, body)
}
~~~  
3. Access logging -
4. Testing -
5. Service hosting -

Resource versioning is imlemented in the timeseries package. Package level access logging is supported via integration with the core.Access package.

Applications that want to use example-domain functionality can integrate directly, by calling the package's Get or Post, or access the functionality hosted in another service. Hosting example-domain packages only requires registering a ServMux handler and pattern, which are both defined in the package.go file. All of the testing, including the Http handler, is automated, in process, and in the package. Additional testing in a service host is not required. This allows the packages to be deployed in multiple hosts, providing flexibility when creating new functionality. New services can utilize existing services, or integrate directly with the packaged functionality. 

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



