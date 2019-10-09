package main

import (
	"encoding/json"
	"fmt"
)

type Result struct {
	Scores []int
}

func main() {
	text := "{\"Scores\":[100,200,300]}"
	bytes := []byte(text)

	var result Result
	json.Unmarshal(bytes, &result)

	fmt.Printf("%v \n", result.Scores)
	for _, value := range result.Scores {
		fmt.Println(value)
	}
}
