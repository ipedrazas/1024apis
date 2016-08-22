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
)

func random(min, max int) int {
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

// generate dependencies.
// Avoid dependening on itself (block int)
// numServices: how many services do we have in total.
// maxDependencies: how many dependencies we want to have as max.
func generateDependencies(numServices int, maxDependencies int, block int) []string {
	numDep := random(0, maxDependencies)
	deps := make([]string, numDep)
	for idx := range deps {
		dep := block
		for dep == block {
			dep = random(1, numServices)
		}
		deps[idx] = fmt.Sprintf("ms-%04d", dep)
	}

	return deps
}

func createService(i int, total int, maxDep int, baseDir string, baseTemplate string) Deployment {
	// generate random json doc
	jsonObj := generateJSON()
	dependencies := generateDependencies(total, maxDep, i)
	b64json := base64.StdEncoding.EncodeToString(jsonObj)
	name := fmt.Sprintf("ms-%04d", i)
	return Deployment{
		Index:        i,
		JSONBody:     b64json,
		Name:         name,
		Dependencies: dependencies,
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

func generateMatrix(services []Deployment, baseDir string) {
	js, err := json.Marshal(services)
	check(err)
	err = ioutil.WriteFile(baseDir+"/matrix.json", js, 0644)
	check(err)
}
