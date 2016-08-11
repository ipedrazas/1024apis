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

// exists returns whether the given file or directory exists or not
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func main() {
	rand.Seed(time.Now().Unix())
	numRes := flag.Int("n", 1, "Number of resources created")
	baseTemplate := flag.String("t", "tmpl/dep.template", "Template used")
	baseDir := flag.String("d", "runtest", "Target dir")
	flag.Parse()

	if _, err := os.Stat(*baseDir); os.IsNotExist(err) {
		os.MkdirAll(*baseDir, 0755)
	}

	for i := 1; i < *numRes+1; i++ {
		createService(i, *baseDir, *baseTemplate)
	}

	// merge.

}
