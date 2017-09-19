package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func() {
		defer func() {
			fmt.Println("step1")
		}()
		go func() {
			fmt.Println("step2")
		}()
		runtime.Goexit() // 仍然不影响defer的执行, 并且不会引起panic
	}()
	fmt.Println("step 0")
	runtime.Goexit() // 会等待所有的Gorountine执行完
	// <-time.After(time.Second * 3)

	// output:
	// step1
	// step 0
	// step2
	// fatal error: no goroutines (main called runtime.Goexit) - deadlock!

}
