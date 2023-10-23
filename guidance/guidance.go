package guidance

// Guidance
// Examples:
//  1. Apply rate limiting before routing for a given controller, or a group of controllers
type Guidance struct {
	Agent       string
	Description string
	// Applies to agent activities
}

// Constraint - applies to agent activities
// Examples:
//  1. No re-routing during the following time duration
type Constraint struct {
	Agent string
}
