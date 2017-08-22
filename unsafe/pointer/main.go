package main

import (
	"fmt"
	"unsafe"
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

}
