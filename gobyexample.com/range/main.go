package main

import (
	"fmt"
)

func main() {
	// 遍历slice
	nums := []int{1, 2, 3}
	// range可以看做迭代器
	for i, num := range nums {
		fmt.Println("index:", i, " num:", num)
	}

	var n int
	for _, num := range nums {
		n += num
	}
	fmt.Println("total:", n)

	// 遍历map
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Println("k:", k, ";v:", v)
	}
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// 遍历字符串 => 输出c的为 rune 类型
	for i, c := range "go" {
		fmt.Println("i:", i, ";c:", c)
	}
}
