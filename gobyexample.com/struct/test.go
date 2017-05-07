package main

import "fmt"

func init() {
	fmt.Println("init")
}

func test() {
	fmt.Println("test")
}

func Test() {
	fmt.Println("can be import")
}

type Demo struct {
	s []int
}

func (d Demo) Add(i int) {
	d.s = append(d.s, i)
}

func (d *Demo) PAdd(i int) {
	d.s = append(d.s, i)
}

type IntStruct int

func (is *IntStruct) String() {
	fmt.Println(*is)
}

type DemoFunc func(s string)

func (fn *DemoFunc) DemoFuncFunc(s string) {
	(*fn)(s)
}

// 动态更改fn
func (fn *DemoFunc) Override(f func(s string)) {
	*fn = f
}
