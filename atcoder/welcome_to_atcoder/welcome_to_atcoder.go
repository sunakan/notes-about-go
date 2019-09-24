package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a, b, c int
	var s string
	fmt.Scan(&a)
	fmt.Scan(&b, &c)
	fmt.Scan(&s)
	var ans = solve(a, b, c, s)
	fmt.Println(ans)
}

func solve(a, b, c int, s string) string {
	var sum = strconv.Itoa(a + b + c)
	return sum + " " + s
}
