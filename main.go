package main

import (
	"flag"
	"fmt"
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
	// first arg is the program itself
	if len(os.Args) == 1 {
		fmt.Println("1024apis, a little tool to simulate microservices systems")
		os.Exit(0)
	}
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
