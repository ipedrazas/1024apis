package main

import "strings"

func toNeo_test(t *T.testing) {
	name := "ms-0009"
	neoName := toNeo(name)
	if strings.Contains(neoName, "-") {

	}
}
