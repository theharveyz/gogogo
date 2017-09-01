package main

import (
	"fmt"
)

type Si interface {
	Demo() int
}
type St struct{}

func (this *St) Demo() int {
	return 0
}

func test() *St {
	return nil
}

func testReturnInterface() Si {
	return nil
}

type People interface {
	Show()
}

type Student struct{}

func (stu *Student) Show() {

}

func live() People {
	var stu *Student
	return stu
}

func main() {
	a := test()
	fmt.Println("a is nil", a == nil)

	b := testReturnInterface()
	fmt.Println("b is nil", b == nil)

	var c Si
	fmt.Println("c is nil", c == nil)
	c = test()
	// 即：有类型的nil，转换为interface类型时，interface要对比:
	// 1. *itab的*_type指针是否是nil,
	// 2. data的*_data指针为nil
	fmt.Println("c is nil", c == nil) // 返回false，很显然此interface的*_type指针不是nil

	// interface比较的是data是否是nil
	if live() == nil {
		fmt.Println("AAAAAAA")
	} else {
		// 执行到这里
		fmt.Println("BBBBBBB")
	}
	var s *Student
	if s == nil {
		fmt.Println("a is nil")
	}
}
