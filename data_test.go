package main

import (
	"encoding/json"
	"testing"
)

func TestRandom(t *testing.T) {
	val := random(0, 10)
	if val < 0 || val > 10 {
		t.Error("val should be between 0 and 10")
	}
}

func TestGenerateString(t *testing.T) {
	val := generateString()
	if len(val) < 3 {
		t.Error("String is too short")
	}
	if len(val) > 10 {
		t.Error("String is too long")
	}
}

func TestGenerateJson(t *testing.T) {
	myjson := generateJSON()
	var js map[string]interface{}
	if json.Unmarshal([]byte(myjson), &js) != nil {
		t.Error("Not valid Json")
	}
}

func TestIntInSlice(t *testing.T) {

	b := []int{1, 2, 3, 4, 5}
	if !intInSlice(1, b) {
		t.Error("Should be in the array")
	}
	if intInSlice(0, b) {
		t.Error("Not in the array")
	}
}

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

func TestAddToSlice(t *testing.T) {
	sl := []int{1, 2, 3, 4}
	sl = addtoSlice(5, sl)
	if !intInSlice(5, sl) {
		t.Error("Int not added to Slice")
	}
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
