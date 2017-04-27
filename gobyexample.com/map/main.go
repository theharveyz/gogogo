package main

import (
	"fmt"
)

func main() {
	// map声明, map底层结构为hashtable, 并未用红黑树
	m := make(map[string]int)

	m["k1"] = 1
	m["k2"] = 2

	fmt.Println(m)
	// 元素个数
	fmt.Println("len:", len(m))

	// 访问不存在的key
	err, value := m["k3"]          // 返回是否存在
	fmt.Println("k3:", value, err) // value 为false, 不为nil, err为0

	// 声明并赋值
	m2 := map[string]int{"foor": 1} // 大括号
	fmt.Println(m2)
}
