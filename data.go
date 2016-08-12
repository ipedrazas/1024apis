package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"html/template"
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

func createService(i int, baseDir string, baseTemplate string) {
	// generate random json doc
	jsonObj := generateJSON()
	b64json := base64.StdEncoding.EncodeToString(jsonObj)
	var buffer bytes.Buffer
	buffer.WriteString(baseDir)
	buffer.WriteString("/")

	name := fmt.Sprintf("ms-%03d", i)
	buffer.WriteString(name)
	d := Deployment{
		Index:    i,
		JSONBody: b64json,
		Name:     name,
	}
	buffer.WriteString(".yaml")
	f, err := os.Create(buffer.String())
	check(err)

	t, _ := template.ParseFiles(baseTemplate) // Parse template file.
	t.Execute(f, d)
}
