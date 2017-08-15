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

func main() {
	a := test()
	fmt.Println("a is nil", a == nil)

	b := testReturnInterface()
	fmt.Println("b is nil", b == nil)

	var c Si
	c = test()
	fmt.Println("c is nil", c == nil) // 返回false， 即：有类型的nil，转换为interface类型时，interface的_data指针不为零值！！
}
