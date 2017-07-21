package main

import (
	"fmt"
)

type Boy int

func af(b int) {
	fmt.Println(b)
}

type BF func(b Boy)

func bf(fn BF) {
	var b Boy = 1
	fn(b)
}

func ff(fn func(b Boy)) {
	var b Boy = 1
	fn(b)
}

func test(b Boy) {
	fmt.Println(b)
}

// 自定义类型转换
func main() {
	bf(test)
	var b BF
	fmt.Println(b) // nil
	b = func(boy Boy) {
		fmt.Println(boy)
	}
	ff(b)

	// var b Boy
	// af(b)
	// 报错： cannot use b (type Boy) as type int in argument to af func(b int)
}
