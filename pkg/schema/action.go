package schema

// Action must be capitalized to be accessible outside of the schema package
type Action struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	// Add your other struct fields here
}
