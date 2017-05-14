package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

// 原子操作

func main() {
	var ops uint64 = 0
	for i := 0; i < 50; i++ {
		go func() {
			for {
				atomic.AddUint64(&ops, 1)
				time.Sleep(time.Microsecond)
			}
		}()
	}

	time.Sleep(time.Second)
	final := atomic.LoadUint64(&ops)
	fmt.Println(final)
}
