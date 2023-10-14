package guidance

type SLO struct {
	// What does this apply to
	Controller string

	// Types of SLOs
	// availability : 99% and 99.999%
	// percentage of traffic : 10% or 10
	// latency percentile: 99/500ms
	Threshold   string // Either percentage of traffic, or latency percentile
	StatusCodes string // For percentage
}

type Update struct {
	Controller string
	Behavior   string // Maybe properties, actions? Can be empty
	Action     string // SQL set statement
}

// Guidance
// Examples:
//  1. Apply rate limiting before routing for a given controller, or a group of controllers
type Guidance struct {
	Agent string
	// Applies to agent activities
}

// Constraint - applies to agent activities
// Examples:
//  1. No re-routing during the following time duration
type Constraint struct {
	Agent string
}
