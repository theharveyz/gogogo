package main

import (
	"fmt"
	"unsafe"

	"./lib"
)

func main() {
	a := "hello"
	var b []byte

	b = *(*[]byte)(unsafe.Pointer(&a))
	fmt.Println("b:", b)
	fmt.Println("a:", a)

	a = *(*string)(unsafe.Pointer(&b))
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("a:", uintptr(unsafe.Pointer(&a)))
	fmt.Println("b:", &b)
	fmt.Println("a:", &a)

	// 获取私有变量的方式
	test := lib.Test{
		Y: 200,
	}
	demo := (*struct{ y int })(unsafe.Pointer(&test))
	demo.y = 100
	fmt.Println("demo:", demo) // demo: &{100}
	fmt.Println("test:", test) // test: {100 200}
}
