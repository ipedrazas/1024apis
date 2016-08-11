package main

import (
	"flag"
	"math/rand"
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
	flag.Parse()

	for i := 1; i < *numRes+1; i++ {
		createService(i, *baseDir, *baseTemplate)
	}

	// merge.

}
