package main

import (
	"fmt"
)

type demo struct {
	S string
}

func new() *demo {
	return &demo{"ptr"}
}

func new1() demo {
	return demo{"obj"}
}

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

	demo := new()
	fmt.Println(demo.S)
	demo2 := demo
	fmt.Println(demo2)
	demo2.S = "demo2"
	fmt.Println(demo.S)
	fmt.Println(demo2.S)
	demo2 = new()
	fmt.Println(demo.S)
	fmt.Println(demo2.S)

	demo3 := new1()
	demo4 := demo3
	fmt.Println("demo3:", demo3.S)
	fmt.Println("demo4:", demo4.S)
	demo4.S = "demo4"
	fmt.Println("demo3:", demo3.S)

	a, c := 1, 2
	b := &a
	fmt.Println(b)
	b = &c
	fmt.Println(b)
	fmt.Println(a)
}
