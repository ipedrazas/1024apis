package main

import (
	"flag"
	"math/rand"
	"os"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	numRes := flag.Int("n", 1, "Number of resources created")
	baseTemplate := flag.String("t", "tmpl/dep.template", "Template used")
	baseDir := flag.String("d", "runtest", "Target dir")
	maxDep := flag.Int("m", 5, "Maximum number of dependencies per service")
	flag.Parse()

	if _, err := os.Stat(*baseDir); os.IsNotExist(err) {
		os.MkdirAll(*baseDir, 0755)
	}

	services := make([]Deployment, *numRes)
	for i := 1; i < *numRes+1; i++ {
		srv := createService(i, *numRes, *maxDep)
		services[i-1] = srv
	}

	generateMatrix(services, *baseDir)
	writeServicesToYaml(services, *baseDir, *baseTemplate)

}
