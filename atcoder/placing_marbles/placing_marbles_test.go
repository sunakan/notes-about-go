package main

import (
	"testing"
)

func TestProduct1(t *testing.T) {
	var ans = solve("101")
	if ans != 2 {
		t.Error("Expected 2, got ", ans)
	}
}

func TestProduct2(t *testing.T) {
	var ans = solve("000")
	if ans != 0 {
		t.Error("Expected 1, got ", ans)
	}
}
