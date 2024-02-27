# example-domain

Provides common functionality utilized by an AI agent managing service resiliency. The packages, activity, slo, and timeseries, provide uniform interfaces in package.go.

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

Package service is the resource that handles incoming HTTP requests. There are seperate handlers for each package plus a search endpoint. Testing is implemented via the uniform HttpHandler interface in each handler.

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
   

Example-domain packages also address the following functional areas:
1. Resource versioning - the timeseries package supports versioning via seperate sub packages, with the exchange functions in package.go.
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
3. Access logging - integration with core/access package.
~~~
defer access.LogDeferred(access.InternalTraffic, access.NewRequest(h, http.MethodGet, getLoc), -1, "", access.NewStatusCodeClosure(&status))()
return getHandler[runtime.LogError](nil, h, u)
~~~
4. Messaging exchange initialization and message handling.
~~~
func init() {
    status := exchange.Register(exchange.NewMailbox(PkgPath, false))
    if status.OK() {
	agent, status = exchange.NewAgent(PkgPath, messageHandler, nil, nil)
    }
    if !status.OK() {
	fmt.Printf("init() failure: [%v]\n", PkgPath)
    }
    agent.Run()
}

func messageHandler(msg core.Message) {
    start := time.Now()
    switch msg.Event {
    case core.StartupEvent:
	core.SendReply(msg, runtime.NewStatusOK().SetDuration(time.Since(start)))
    case core.ShutdownEvent:
    case core.PingEvent:
	core.SendReply(msg, runtime.NewStatusOK().SetDuration(time.Since(start)))
    }
}
~~~
5. Testing - all testing, including the Http handler, is automated, in process, and in the package. Additional testing in a service host is not needed. The http_test.go file utilizes functionality in the core/http2/http2test package, with all test requests and responses deserialized from disk.

6. Service integration - Applications that want to use example-domain functionality can integrate directly, by calling the package's Get or Post, or access the functionality hosted in another service. Hosting example-domain packages only requires registering a ServMux handler and pattern, which are both defined in the package.go file. Packages can be deployed in multiple hosts, providing flexibility when creating new functionality, as new services can utilize existing services, or integrate directly with the packaged functionality. 


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



