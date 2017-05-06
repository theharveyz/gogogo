package main

import "fmt"

type person struct {
	name                string
	age, weight, height int
}

func (p *person) bmiP() float64 {
	// 使用指针作为receiver:
	// 1. 可以改变内部属性
	// 2. 避免对象拷贝
	p.age = 18 // 这里改变了对象的属性
	if p.height <= 0 {
		return float64(0)
	}
	return float64(p.weight / (p.height * p.height))
}

func (p person) bmi() float64 {
	p.age = 18 //
	if p.height <= 0 {
		return float64(0)
	}
	return float64(p.weight / (p.height * p.height))
}

func main() {
	fmt.Println(person{name: "harveyz", age: 25, weight: 175})
	fmt.Println(&person{"harveyz", 25, 150, 175})
	s := person{name: "harveyz"}
	fmt.Println(s.name)
	fmt.Println(s.age)

	sp := &s // &{harveyz 25 175} --- 并非是取地址符, 而是引用符号
	fmt.Println(sp.name)
	sp.name = "harveyz1"
	fmt.Println(s.name == sp.name) // name 也跟着改变 => true

	harveyz := person{"harveyz", 25, 150, 175}
	fmt.Println("bmi:", harveyz.bmi())
	fmt.Println("age是否改变:", harveyz.age) // 否 => age是否改变: 25
	fmt.Println("bmiP:", harveyz.bmiP())
	fmt.Println("age是否改变:", harveyz.age) // age是否改变: 18 => 是
	fmt.Println("bmiP:", (&harveyz).bmiP())

	// same module
	test()
	Test()

	// 测试receiver的引用问题
	demo := &Demo{}
	demo.Add(1)
	fmt.Println(demo.s) //  [] ====> Add里的 receiver为重新分配内存, 与demo不存在引用关系

	demo.PAdd(1)
	demo.PAdd(1)
	fmt.Println(demo.s) // [1 1] ===> 与receiver存在引用关系

	// 任何类型都可以设置方法
	var is IntStruct = 1111
	fmt.Println(is.String())

}
