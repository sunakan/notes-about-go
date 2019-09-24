package main

import (
	"testing"
)

func TestProduct1(t *testing.T) {
	var ans = solve(3, 4)
	if ans != "Even" {
		t.Error("Expected 'Even', got ", ans)
	}
}

func TestProduct2(t *testing.T) {
	var ans = solve(1, 21)
		if ans != "Odd" {
		t.Error("Expected 'Odd', got ", ans)
	}
}
