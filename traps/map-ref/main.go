package main

import "fmt"

func main() {
	dict := map[int]int{
		1: 1,
		2: 2,
	}
	fmt.Printf("%p\n", &dict) // 0xc042004028

	dict = map[int]int{
		3: 1,
		4: 2,
	} // 0xc042004028
	fmt.Printf("%p\n", &dict)
	d := dict
	fmt.Printf("%p\n", &d) // 0xc042052028

	i := 1
	fmt.Printf("%p\n", &i) // 0xc04200c2b8
	i = 2
	fmt.Printf("%p\n", &i) // 0xc04200c2b8

	// -- output:
	// 重新赋值，不会更改变量地址（不同于指针）
	// 不是严格意义上的 引用 类型

}
