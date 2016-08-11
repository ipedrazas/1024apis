package main

// A Deployment represents a Kubernetes Deployment
// artifact
type Deployment struct {
	Index    int
	JSONBody string
	Name     string
}
