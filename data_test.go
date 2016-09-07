package main

import "testing"

func TestAllowDependency(t *testing.T) {

	dep := Dependencies{
		List:      []int{1, 2, 3},
		Forbidden: []int{1, 6},
		Index:     9,
	}

	if !allowedDependency(12, dep) {
		t.Error("Dependency not allowed")
	}
	if allowedDependency(6, dep) {
		t.Error("Dependency not allowed")
	}
}

func TestCreateService(t *testing.T) {

	dep := Dependencies{
		List:      []int{1, 2, 3},
		Forbidden: []int{1, 6},
		Index:     9,
	}

	srv := createService(dep)
	if srv.Index != 9 {
		t.Error("Index is wrong")
	}
	if srv.Name != "ms-0009" {
		t.Error("Name not set properly")
	}
	// js, _ := json.Marshal(srv)
	// fmt.Println(string(js))
}

func TestDependencyMap(t *testing.T) {
	dep := createDependencyMap(10, 3)
	// DependencyMap has a zero
	if len(dep) != 11 {
		t.Error("DependencyMap size is wrong")
	}
	// js, _ := json.Marshal(dep)
	// fmt.Println(string(js))
}

func TestCreateServices(t *testing.T) {
	total := 10
	maxDependencies := 3
	services := createServices(total, maxDependencies)
	if len(services) != 10 {
		t.Error("Services sizes is not right")
	}
}
