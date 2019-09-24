package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	var ans = solve(a, b)
	fmt.Println(ans)
}

func solve(a, b int) string {
	if a*b % 2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}
