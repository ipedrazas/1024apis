package main

import (
	"encoding/base64"
	"fmt"
)

func intInSlice(index int, list []int) bool {
	for _, v := range list {
		if v == index {
			return true
		}
	}
	return false
}

func createServices(num int, maxDependencies int) []Deployment {
	deps := createDependencyMap(num, maxDependencies)
	services := make([]Deployment, len(deps)-1)
	for i, dep := range deps {
		if i > 0 {
			services[i-1] = createService(dep)
		}
	}
	return services
}

func createDependencyMap(num int, maxDependencies int) []Dependencies {
	dep := make([]Dependencies, num+1)
	// initialise
	for i := 1; i < num+1; i++ {
		numDep := random(0, maxDependencies)
		dep[i] = Dependencies{
			Index:     i,
			List:      make([]int, numDep),
			Forbidden: []int{i},
		}
	}
	dep = createDependencies(maxDependencies, dep)

	return dep
}

func createDependencies(maxDependencies int, deps []Dependencies) []Dependencies {

	for index, dependency := range deps {
		if index > 0 {
			// let's populate dependencies
			for idx := range dependency.List {
				dep := index
				for !allowedDependency(dep, dependency) {
					dep = random(1, len(deps))
				}
				dependency.List[idx] = dep
				target := deps[dep]
				if !intInSlice(index, target.Forbidden) {
					target.Forbidden = walkDependencies(dependency, deps, target.Forbidden)
					deps[dep] = target
				}
			}
			deps[index] = dependency
		}
	}
	return deps
}

func createService(dependencies Dependencies) Deployment {
	// generate random json doc
	jsonObj := generateJSON()
	b64json := base64.StdEncoding.EncodeToString(jsonObj)
	deps := make([]string, len(dependencies.List))
	depNames := make([]string, len(dependencies.List))
	for idx, dependecy := range dependencies.List {
		deps[idx] = fmt.Sprintf("http://ms-%04d-srv:5000/srv%d", dependecy, dependecy)
		depNames[idx] = fmt.Sprintf("ms-%04d", dependecy)
	}

	return Deployment{
		Index:             dependencies.Index,
		JSONBody:          b64json,
		Name:              fmt.Sprintf("ms-%04d", dependencies.Index),
		DependencyIndexes: dependencies.List,
		Dependencies:      deps,
		DependencyNames:   depNames,
		ForbiddenIndexes:  dependencies.Forbidden,
	}
}

func allowedDependency(testDep int, dependency Dependencies) bool {
	// check that the dep is not in blocked
	if intInSlice(testDep, dependency.Forbidden) {
		return false
	}
	return true
}

func walkDependencies(dep Dependencies, allDeps []Dependencies, blocked []int) []int {
	if !intInSlice(dep.Index, blocked) {
		// first we add the current dependency
		blocked = addtoSlice(dep.Index, blocked)
		for _, idxDep := range dep.Forbidden {
			blocked = walkDependencies(allDeps[idxDep], allDeps, blocked)
		}
	}
	return blocked
}
