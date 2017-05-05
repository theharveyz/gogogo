package main

import (
	"fmt"
	"time"
)

func main() {
	// if else块 与 块级变量作用域
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num > 8 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	// fmt.Println(num) ---- num is undefined

	// switch case无需break,默认break.switch甚至不需要后面带着变量
	now := time.Now()
	switch {
	case now.Hour() < 12:
		fmt.Println(now.Hour(), ";It is before noon")
	default:
		fmt.Println("It is after noon")
	}
	// switch带着赋值
	switch now.Weekday() {
	case time.Saturday, time.Sunday: // 多个匹配值
		fmt.Println("It is a weekend")
	default:
		fmt.Println("It is an workday")
	}

	// switch断言语句
	// 只能在switch语句中使用
	var day interface{} = now.Day()

	switch x := day.(type) { // 局部变量声明了,也必须要使用才行
	case int:
		fmt.Println("int", x)
	case bool:
		fmt.Println("boolean", x)
	case string:
		fmt.Println("string", x)
	}
}
