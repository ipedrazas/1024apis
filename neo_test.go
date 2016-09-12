package main

import (
	"strings"
	"testing"
)

func TestToNeo(t *testing.T) {
	name := "ms-0009"
	neoName := toNeo(name)
	if strings.Contains(neoName, "-") {

	}
}
