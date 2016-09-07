package main

import (
	"encoding/json"
	"testing"
)

func TestRandom(t *testing.T) {
	val := random(0, 10)
	if val < 0 || val > 10 {
		t.Error("val should be between 0 and 10")
	}
}

func TestGenerateString(t *testing.T) {
	val := generateString()
	if len(val) < 3 {
		t.Error("String is too short")
	}
	if len(val) > 10 {
		t.Error("String is too long")
	}
}

func TestGenerateJson(t *testing.T) {
	myjson := generateJSON()
	var js map[string]interface{}
	if json.Unmarshal([]byte(myjson), &js) != nil {
		t.Error("Not valid Json")
	}
}

func TestIntInSlice(t *testing.T) {

	b := []int{1, 2, 3, 4, 5}
	if !intInSlice(1, b) {
		t.Error("Should be in the array")
	}
	if intInSlice(0, b) {
		t.Error("Not in the array")
	}
}

func TestAddToSlice(t *testing.T) {
	sl := []int{1, 2, 3, 4}
	sl = addtoSlice(5, sl)
	if !intInSlice(5, sl) {
		t.Error("Int not added to Slice")
	}
}
