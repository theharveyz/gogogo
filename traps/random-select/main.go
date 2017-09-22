package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1)
	go func(c <-chan int) {
		for i := range c {
			fmt.Println(i)
		}
	}(ch)

	go func(c chan<- int) {
		for i := 0; i < 10; i++ {
			// 注意:
			// 1. 同switch case一样, 可以省略return
			// 2. select会随机选择一个case执行
			// 3. select本身是一个随机选择!!!  而不是 channel receiver监听独有!!!
			select {
			case c <- i * 1:
			case c <- i * 2:
			}
		}
		close(c)
	}(ch)

	// output:
	// 0
	// 2
	// 2
	// 3
	// 4
	// 10
	// 12
	// 14
	// 16
	// 9
	<-time.After(time.Second * 3)
}
