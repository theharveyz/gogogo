package main

import (
	"fmt"
	"os"
)

func main() {
	ch := make(chan int)
	ch2 := make(chan bool)
	go func() {
		close(ch)
	}()
	// 关闭channel select case也可以执行
	for {
		select {
		case <-ch:
			fmt.Println("channel1 closed")
			ch2 <- true // fatal error: all goroutines are asleep - deadlock!
		case <-ch2:
			fmt.Println("chan2 closed")
			os.Exit(1)
		}
	}

	// output:
	// channel closed
	// fatal error: all goroutines are asleep - deadlock!
}
