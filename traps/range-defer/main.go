package main

import "fmt"
import "errors"
import "time"

// 多defer与panic
func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	go func() {
		defer func() {
			fmt.Println("defer 1")
		}()
		defer func() {
			defer func() {
				if err := recover(); err != nil {
					fmt.Println(err)
				}
				panic(errors.New("oh fxxk!"))
			}()
			fmt.Println("defer 2")
			panic(errors.New("oh no!"))
		}()
		defer func() {
			fmt.Println("defer 3")
		}()

		for i := 0; i < 3; i++ {
			defer func(i int) {
				fmt.Println("range defer:", i)
			}(i)
		}
	}()
	// panic(errors.New("oh fxxk!"))

	<-time.After(time.Second * 3)

	// output:
	// range defer: 2
	// range defer: 1
	// range defer: 0
	// defer 3
	// defer 2
	// oh no!
	// defer 1

	// 结论:
	//	0. defer 可以嵌套在range和for循环中
	// 	1. panic不影响同级defer的FIFO执行
	//  2. 函数的defer中的 recover 不能捕获该函数 goroutine的panic
}
