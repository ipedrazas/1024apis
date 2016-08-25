package main

// A Deployment represents a Kubernetes Deployment
// artifact
type Deployment struct {
	Index             int
	JSONBody          string
	Name              string   `json:"name"`
	Dependencies      []string `json:"dependencyUrls"`
	DependencyNames   []string `json:"dependencies"`
	DependencyIndexes []int
	ForbiddenIndexes  []int
}

// Dependencies represents the relationship/dependencies
// between different services
type Dependencies struct {
	List      []int
	Forbidden []int
	Index     int
}
