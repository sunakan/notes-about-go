package main

import (
	"encoding/json"
	"fmt"
)

type Human struct {
	Name string
	Age  int
	Weight int
}

func main() {
	human := Human {
		Name: "Taro",
		Age:  0,
		Weight: 1000,
	}

	bytes, _ := json.Marshal(human)
	string := string(bytes)
	fmt.Println(string)
}
