package timeseries

import "time"

// Entry - timeseries struct
type Entry struct {
	Traffic        string
	Start          time.Time
	Duration       time.Duration
	ControllerName string

	Region     string
	Zone       string
	SubZone    string
	Service    string
	InstanceId string

	Uri string // {scheme}://{host}/{path}
	// How to determine if uri is primary or secondary, should use traffic percentage
	Method      string
	StatusCode  int32
	StatusFlags string

	// Needed to verify client controller configuration matches configuration in cloud
	// Can this be replaced with a periodic audit?
	Timeout           int32
	RateLimit         float64
	RateBurst         int32
	TrafficPercentage int32
	// Proxy methods  ??

}
