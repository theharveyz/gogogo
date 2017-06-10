package main

import (
	"fmt"
)

func main() {
	s1 := make([]int, 2, 10)
	s2 := append(s1, 1)
	fmt.Println(s1) // [0, 0]
	fmt.Println(s2) // [0, 0, 1]
}
