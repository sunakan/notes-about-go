package main

import (
	"encoding/json"
	"fmt"
)

type Human2 struct {
	Name string
	Age  int
}

func main() {
	text := "[{\"name\":\"Taro\",\"Age\":20},{\"Name\":\"Jiro\",\"Age\":15}]"
	bytes := []byte(text)
	var humans []Human2
	json.Unmarshal(bytes, &humans)

	for i := range humans {
		fmt.Printf("【氏名】%v  【年齢】%v", humans[i].Name, humans[i].Age)
		fmt.Println()
	}

}
