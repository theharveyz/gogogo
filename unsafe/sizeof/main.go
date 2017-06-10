package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// unsafe.Sizeof 返回的不是实际变量占用的字节数 --- 注意为字节数
	// 而是按照变量的"top level"内存计算: 比如 string类型存储的是指针等等

	var a int
	fmt.Println(unsafe.Sizeof(a)) // 8
	var s struct{}
	fmt.Println(unsafe.Sizeof(s)) // 0
	var n interface{}
	fmt.Println(unsafe.Sizeof(n)) // 16

	var np *interface{}
	fmt.Println(unsafe.Sizeof(np)) // 8 指针类型

	var b string
	fmt.Println(unsafe.Sizeof(b))  // 16 => pointer + len + cap
	fmt.Println(unsafe.Sizeof(&b)) // 8 ==> 指针类型占用8个字节

	var nos struct{ a string }
	fmt.Println(unsafe.Sizeof(&nos)) // 8 ==> 指针类型占用8个字节
	fmt.Println(unsafe.Sizeof(nos))  // 16 ==> a 类型占用的字节数

	var arr [1]int
	fmt.Println(unsafe.Sizeof(arr)) // 8
	var arr100 [100]int
	fmt.Println(unsafe.Sizeof(arr100)) // 8 * 100 // 连续的空间(指针)

	var sl []int
	fmt.Println(unsafe.Sizeof(sl)) // 24 ==> 首地址 + len + cap
	ls := append(sl, 10, 10)
	fmt.Println(unsafe.Sizeof(ls)) // 24 ==> 首地址(底层对应数组的首地址) + len + cap

	var bo bool
	fmt.Println(unsafe.Sizeof(bo)) // 1 ==> true / false 比较特殊, 自身就是值!! 只占 1直接 8bit

	var u8 uint8
	fmt.Println(unsafe.Sizeof(u8))  // 1 ==> int类型的值就是其自身
	fmt.Println(unsafe.Sizeof(&u8)) // 8 ==> pointer

	var i64 int64
	fmt.Println(unsafe.Sizeof(i64)) // 8

	var r rune
	fmt.Println(unsafe.Sizeof(r)) // 4

}
