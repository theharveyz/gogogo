package main

import (
	"fmt"
)

type Parent struct{}
type OtherParent struct{}

func (Parent) Test() {
	fmt.Println("parent test")
}
func (Parent) Test2() {
	fmt.Println("parent test2")
}
func (OtherParent) Other() {
	fmt.Println("parent other")
}

type Child struct {
	Parent
	OtherParent OtherParent
}

func (c Child) Test() {
	fmt.Println("child test")
}

type Demoer interface {
	Test()
}
type Demoer2 interface {
	Demoer
}

func main() {
	var c Child
	// implement
	c.Test()
	c.Parent.Test()
	// child test
	// parent test

	// 匿名属性
	c.Test2() // 两者等价
	c.Parent.Test2()
	// parent test2
	// parent test2

	// 显式声明属性方式
	c.OtherParent.Other()
	// c.Other()  --- 则这种写法则不可以

	var demo Demoer2
	demo = c
	demo.Test()
}
