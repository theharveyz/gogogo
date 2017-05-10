package main

import (
	"fmt"
)

func main() {
	fmt.Println("channel关闭后仍然可以通过range获取剩余值")
	queue := make(chan string, 3)
	queue <- "one"
	queue <- "two"
	close(queue)
	i, more := <-queue // 这里同样可以读取
	// queue <- "hello" // 不能写入
	fmt.Println(i, more)
	for elem := range queue {
		fmt.Println(elem)
	}
}
