package main

import (
	"fmt"
)

func main() {
	s := "请求"
	b := "hello"
	sb := s + b
	fmt.Println(sb)
	sbb := []byte(sb)
	fmt.Println(sbb)
	fmt.Println(string(sbb))
	s = string(sbb)
	fmt.Println(s)
}
