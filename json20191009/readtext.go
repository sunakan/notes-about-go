package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	bytes, err := ioutil.ReadFile("./sample.json")
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Println("------------")
	fmt.Println(string(bytes))
	fmt.Println("------------")
}
