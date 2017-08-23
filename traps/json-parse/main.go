package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Enum struct {
	Key string
}
type Test struct {
	Enum
}

type Test2 struct {
	Enum Enum
}

func main() {
	file, _ := ioutil.ReadFile("./test.json")
	var t Test
	json.Unmarshal(file, &t)
	fmt.Println(t.Enum) // {} --- 这种情况则解析不了

	var t2 Test2
	json.Unmarshal(file, &t2)
	fmt.Println(t2.Enum.Key) // output: enum --- 这种写法则可以解析
}
