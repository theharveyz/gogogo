package main

import (
	"fmt"
)

func test() int {
	i := 1
	defer func() {
		i++
	}()
	// 返回的将是 1, defer执行之前return的变量先发生一次拷贝
	return i
}

func main() {
	fmt.Println(test())
}
