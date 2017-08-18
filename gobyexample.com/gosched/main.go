package main

import (
	"fmt"
	"runtime"
	"sync"
)

var counter int = 0

func Count(lock *sync.Mutex) {
	lock.Lock()
	defer lock.Unlock()
	// 加加(a = a + 1)不是个原子性操作, 同时写入时可能并不能保证递增
	// 比如:
	// counter = 5
	// counter = 5
	counter++
	fmt.Println("counter =", counter)
}

func main() {

	lock := &sync.Mutex{}

	for i := 0; i < 10; i++ {
		go Count(lock)
	}

	for {
		lock.Lock()
		if counter >= 10 {
			break
		}
		lock.Unlock()

		runtime.Gosched() // 出让时间片, 让其他goroutine执行, 然后继续执行当前for循环以下的操作
		fmt.Println(counter)
		fmt.Println("hahahaha") // 输出的次数不固定, 依次可以看出来 Gosched并非是等其他goroutine都执行完后再切回
		// -- output:
		// hahahaha
		// counter = 3
		// counter = 4
		// counter = 5
		// counter = 6
		// counter = 7
		// 5
		// counter = 8
		// hahahaha
		// counter = 9
		// counter = 10

	}
}
