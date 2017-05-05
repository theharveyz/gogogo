package main

import (
	"fmt"
)

// 闭包作为返回值时:
func intSeq() func(a, b int) (int, int) {
	i := 0
	return func(a, b int) (int, int) {
		i++
		return a + b, i
	}
}
func main() {
	fmt.Println(intSeq()(1, 2))
	fmt.Println(intSeq()(1, 2))
	demo := intSeq()
	fmt.Println(demo(1, 2)) // 3 1
	fmt.Println(demo(1, 2)) // 3 2
	fmt.Println(demo(1, 2)) // 3 3
}
