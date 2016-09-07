package main

import (
	"bytes"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

func generateMatrix(services []Deployment, baseDir string) {
	js, err := json.Marshal(services)
	checkError("Cannot mashal services", err)
	err = ioutil.WriteFile(baseDir+"/matrix.json", js, 0644)
	checkError("Cannot write matrix.json", err)
}

func writeServicesToYaml(services []Deployment, baseDir string, baseTemplate string) {
	for _, d := range services {
		var buffer bytes.Buffer
		buffer.WriteString(baseDir)
		buffer.WriteString("/")
		buffer.WriteString(d.Name)
		buffer.WriteString(".yaml")
		f, err := os.Create(buffer.String())
		checkError("Cannot create file "+buffer.String(), err)
		t, _ := template.ParseFiles(baseTemplate) // Parse template file.
		t.Execute(f, d)
	}
}
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

func addtoSlice(i int, arr []int) []int {
	length := len(arr)
	arr = append(arr, length+1)
	arr[length] = i
	return arr
}
