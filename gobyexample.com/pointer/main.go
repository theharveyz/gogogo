package main

import (
	"fmt"
)

// 指针的操作
// 传值
func zeroval(i int) {
	i = 0
}

// 传址
func zeroptr(p *int) {
	*p = 0 // 直接解引用, 找到具体地址更改数据, 跟地址有关的变量的值也跟着更改
}

func main() {
	i := 1
	fmt.Println("INIT:", i)
	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer addr:", &i)

	// INIT: 1
	// zeroval: 1
	// zeroptr: 0
	// pointer addr: 0xc42000e280
}
