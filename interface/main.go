package main

import (
	"fmt"
)

type getmetry interface {
	area() float64
}

type rect struct {
	height, weight float64
}

func (r rect) area() float64 {
	return r.height * r.weight
}

func measure(g getmetry) {
	fmt.Println(g.area())
}

func assert(s interface{}) {
	// 类型断言, s必须是接口类型的变量
	t, _ := s.(getmetry)
	fmt.Println(t)

	// type-switch
	switch t := s.(type) { // 只有一个参数, case判断的是type,而不是t, t: {1 100}
	case getmetry:
		fmt.Println("select getmetry:", t)
	}

}

func main() {
	r := rect{height: 1, weight: 100}
	measure(r)
	measure(&r)
	assert(r)
	assert(&r)
	// i := 1
	// i.(int) ---> 报错: i不是接口类型的变量
}
