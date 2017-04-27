package main

import (
	"fmt"
)

func main() {
	s := make([]string, 3, 4) // *arr, len, cap
	fmt.Println("empty slice:", s)
	s[0] = "a"
	s[1] = "b"
	fmt.Println(s, len(s)) // len: 3
	s[2] = "d"
	fmt.Println(len(s))

	// append
	// 大于 cap值,则会
	s = append(s, "e")
	fmt.Println(len(s), "cap:", cap(s))
	s = append(s, "e", "f")
	fmt.Println(len(s), "cap:", cap(s)) // cap 成倍增加
	fmt.Println(s[4])

	s1 := make([]int, 3)
	fmt.Println("len:", len(s1), ";cap:", cap(s1)) // 默认为 初始化len的长度

	// copy ==> 复制数据 --> 引用分离
	c := make([]string, len(s))
	copy(c, s)
	c[1] = "111111"
	fmt.Println(c, s)
	cl3 := c[:3]
	fmt.Println(cl3)

	// 两个slice指向同一个底层数组
	cl3[2] = "cccc"
	fmt.Println(cl3, c)

	// 声明时赋值
	cls := []string{"1", "2", "3"}
	fmt.Println(len(cls), cap(cls)) // 3, 3

	// 多维slice
	twoDS := [][]int{[]int{1}}
	fmt.Println(twoDS, ";len:", len(twoDS))

	twoDS1 := make([][]int, 3)
	for i := 0; i < 3; i++ {
		l := i + 1
		twoDS1[i] = make([]int, l)
		for n := 0; n < l; n++ {
			twoDS1[i][n] = i + n
		}
	}
	fmt.Println(twoDS1)

	// append 超过cap后, 会重新分配一段连续内存,并将原有关联数组copy过去
	one := make([]int, 1)
	one[0] = 1
	two := append(one, 2)
	two[0] = 2
	fmt.Println(one, two) // [1], [2,2]
}
