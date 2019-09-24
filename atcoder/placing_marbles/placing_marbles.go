package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)
	var ans = solve(s)
	fmt.Println(ans)
}

func solve(s string) int {
	var count = 0;
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			count++;
		}
	}
	return count;
}

