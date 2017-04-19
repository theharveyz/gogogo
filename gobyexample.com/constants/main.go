package main

import (
	"fmt"
	"math"
)

const s string = "constant"
const demo = "demo"

func main() {
	demo := "demo-inner"
	const (
		n = 1000
		i = 3e20 / n // 虽然在一个什么块内,但是可以直接使用
	)
	fmt.Println(demo) // 非常量-函数内声明,但是未使用,会编译不通过
	fmt.Println(i)
	fmt.Println(math.Cos(n))
}
