package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

func random(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max-min) + min
}

func generateString() string {
	numL := random(3, 10)
	s := make([]byte, numL)
	for j := 0; j < numL; j++ {
		s[j] = 'a' + byte(rand.Int()%26)
	}
	return string(s)
}

func generateJSON() []byte {
	var buffer bytes.Buffer
	buffer.WriteString("{")
	numAtts := random(1, 10)
	for i := 0; i < numAtts; i++ {
		buffer.WriteString("\"")
		buffer.WriteString(generateString())
		buffer.WriteString("\"")
		buffer.WriteString(":")
		buffer.WriteString("\"")
		buffer.WriteString(generateString())
		buffer.WriteString("\"")
		if i != numAtts-1 {
			buffer.WriteString(",")
		}
	}
	buffer.WriteString("}")

	return buffer.Bytes()
}

func intInSlice(index int, list []int) bool {
	for _, v := range list {
		if v == index {
			return true
		}
	}
	return false
}

func allowedDependency(dependency int, AllDependencies []int, blocked []int) bool {
	return !intInSlice(dependency, blocked) && !intInSlice(dependency, AllDependencies)
}

// generate dependencies.
// Avoid dependening on itself (block int)
// Dependencies cannot repeat
// numServices: how many services do we have in total.
// maxDependencies: how many dependencies we want to have as max.
func generateDependencies(numServices int, maxDependencies int, block []int) []int {
	// makes no sense to have more dependencies than services
	if maxDependencies > numServices {
		maxDependencies = numServices
	}
	numDep := random(0, maxDependencies)
	deps := make([]int, numDep)
	dep := block[0]
	for idx := range deps {
		for !allowedDependency(dep, deps, block) {
			dep = random(1, numServices)
		}
		deps[idx] = dep
		block = append(block, len(block)+1)
		block[len(block)-1] = dep
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

func writeServicesToYaml(services []Deployment, baseDir string, baseTemplate string) {
	for _, d := range services {
		var buffer bytes.Buffer
		buffer.WriteString(baseDir)
		buffer.WriteString("/")
		buffer.WriteString(d.Name)
		buffer.WriteString(".yaml")
		f, err := os.Create(buffer.String())
		check(err)
		t, _ := template.ParseFiles(baseTemplate) // Parse template file.
		t.Execute(f, d)
	}
}

func addtoSlice(i int, arr []int) []int {
	length := len(arr)
	arr = append(arr, length+1)
	arr[length] = i
	return arr
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

func createDependencies(maxDependencies int, deps []Dependencies) []Dependencies {
	// let's calculate how many dependencies this entry has
	for index, dependency := range deps {
		if index > 0 {
			for idx := range dependency.List {
				dep := index
				for !allowedDependency(dep, dependency.List, dependency.Forbidden) {
					dep = random(1, len(deps))
				}
				dependency.List[idx] = dep
				target := deps[dep]
				if !intInSlice(index, target.Forbidden) {
					target.Forbidden = addtoSlice(index, target.Forbidden)
					deps[dep] = target
				}
			}
			deps[index] = dependency
		}
	}
	return deps
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

func generateMatrix(services []Deployment, baseDir string) {
	js, err := json.Marshal(services)
	check(err)
	err = ioutil.WriteFile(baseDir+"/matrix.json", js, 0644)
	check(err)
}
