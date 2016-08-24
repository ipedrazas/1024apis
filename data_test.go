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

	blocked := []int{1, 2, 3, 4}
	allDeps := []int{6, 7, 8}

	if !allowedDependency(12, allDeps, blocked) {
		t.Error("Dependency not allowed")
	}
	if allowedDependency(3, allDeps, blocked) {
		t.Error("Dependency not allowed")
	}
}

func TestGenerateDependency(t *testing.T) {
	maxDeps := 2
	block := []int{3}
	newDep := generateDependencies(3, maxDeps, block)
	if len(newDep) > maxDeps {
		t.Error("Too many Dependencies")
	}
	if intInSlice(block[0], newDep) {
		t.Error("Blocked indexes should not be in dependencies")
	}
}

func TestCreateService(t *testing.T) {
	// index int, total int, maxDep int
	//
	index := 3
	total := 100
	maxDep := 3
	srv := createService(index, total, maxDep)
	if srv.Index != index {
		t.Error("Index is wrong")
	}
	if len(srv.DependencyIndexes) > maxDep {
		t.Error("Too many dependencies")
	}
	// js, _ := json.Marshal(srv)
	// fmt.Println(string(js))
}
