package main

import (
	"fmt"
)

func main() {
	const (
		I       int = 1
		M           = iota
		S           = iota
		a, b, c     = `
		||||||
		\\\\\
		`, "\\\\\\", '#' // 35
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println("go" + "lang")
	fmt.Println(I, M)
	fmt.Println(I + M)
	fmt.Println(true && false)

	// params
	fmt.Println(10 << 1)
	fmt.Println(1 | 1)       // 或运算
	fmt.Println(^2)          // 取反 => 1  一元运算符 -3
	fmt.Println((I + M) ^ 1) // 异或 XOR => 3
	fmt.Println((I + M) ^ 3) // 异或 XOR => 1
	fmt.Println(3 &^ 1)      // AND NOT 位清除运算(类似减法), 将两者都为1的位置为0,其他不变(以前面为主) => 2
	fmt.Println(I / S)       // 0
	fmt.Println(1.0 / S)     // 0.5
}
