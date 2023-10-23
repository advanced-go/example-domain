package objectives

// Guidance
// Examples:
//  1. Apply rate limiting before routing for a given controller, or a group of controllers
type Guidance struct {
	Agent       string
	Description string
	// Applies to agent activities
}

var glist []Guidance

func GetGuidance() []Guidance {
	return glist
}

func AddGuidance(src []Guidance) {
	copy(glist, src)
}

// Constraint - applies to agent activities
// Examples:
//  1. No re-routing during the following time duration
type Constraint struct {
	Agent string
}

var clist []Constraint

func GetConstraint() []Constraint {
	return clist
}

func AddConstraint(src []Constraint) {
	copy(clist, src)
}
