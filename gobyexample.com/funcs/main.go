package main

import (
	"fmt"
)

func plusPlus(args ...int) int {
	var total int
	for _, num := range args {
		total += num
	}
	fmt.Println(total)
	return total
}

func plus(a, b, c int) (int, error) {
	return a + b + c, nil
}

func main() {
	args := []int{1, 2, 3}
	fmt.Println(plus(1, 2, 3))

	// 只有可变参数的函数调用时才能使用 展开式
	// plus函数就不行
	fmt.Println(plusPlus(args...))
	// 使用了可变参数的函数,可以正常方式传参调用
	plusPlus(1, 2, 3)
}
