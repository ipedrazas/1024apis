package main

import (
	"strings"
	"testing"
)

func toNeo_test(t *testing.T) {
	name := "ms-0009"
	neoName := toNeo(name)
	if strings.Contains(neoName, "-") {

	}
}
