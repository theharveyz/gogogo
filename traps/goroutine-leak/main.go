package main

import (
	"runtime"
	"time"
)

func main() {
	ch := make(chan int, 1)
	go func() {
		for i := 0; i < 10; i++ {
			go func() {
				<-ch
			}()
		}
	}()

	// goroutine leak 资源泄露:

	// for循环
	for {
		time.Sleep(time.Second)
		runtime.GC() // 强制垃圾回收
	}
}
