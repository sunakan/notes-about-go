package main

import (
	"testing"
)

func TestProduct1(t *testing.T) {
	var ans = solve(1, 2, 3, "test")
	if ans != "6 test" {
		t.Error("Expected '6 test', got ", ans)
	}
}

func TestProduct2(t *testing.T) {
	var ans = solve(72, 128, 256, "myonmyon")
	if ans != "456 myonmyon" {
		t.Error("Expected '456 myonmyon', got ", ans)
	}
}
