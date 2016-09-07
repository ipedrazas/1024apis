package main

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"time"
)

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
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

	services := createServices(*numRes, *maxDep)

	generateMatrix(services, *baseDir)
	writeServicesToYaml(services, *baseDir, *baseTemplate)
	writeServicesToCSV(services, *baseDir)

}
