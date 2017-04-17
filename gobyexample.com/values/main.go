package main

import (
	"fmt"
)

func main() {
	const (
		I int = 1
		M     = iota
		S     = iota
	)

	fmt.Println("go" + "lang")
	fmt.Println(I, M)
	fmt.Println(I + M)
	fmt.Println(true && false)

	// params
	fmt.Println(10 << 1)
	fmt.Println(1 | 1)       // 或运算
	fmt.Println(^2)          // 取反 => 1
	fmt.Println((I + M) ^ 1) // 异或 XOR => -3
	fmt.Println(3 &^ 1)      // AND NOT 位清除运算, 将两者都为1的位置为0,其他不变 => 2
	fmt.Println(I / S)       // 0
	fmt.Println(1.0 / S)     // 0.5
}
