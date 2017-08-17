package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestWaitGroup(t *testing.T) {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		// Done的定义
		// func (wg *WaitGroup) Done() {
		// 	wg.Add(-1)
		// }
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
			time.Sleep(time.Second * 1)
		}(i)
	}

	wg.Wait() //
	fmt.Println("wait group exit")
}
