package main

import (
	"fmt"
)

func main() {
	s1 := make([]int, 2, 10)
	s2 := append(s1, 1)
	fmt.Println(s1) // [0, 0]
	fmt.Println(s2) // [0, 0, 1]
	s1[1] = 1
	fmt.Println(s2) // [0 1 1]

	// ---- 超过cap --- 重新分配底层array给新的slice
	s1 = make([]int, 2)
	s2 = append(s1, 1)
	s1[1] = 2
	fmt.Println(s1) // [0 2]
	fmt.Println(s2) // [0 0 1]
}
