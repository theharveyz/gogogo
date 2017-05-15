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
	}

	for i := 0; i < 10; i++ {
		// Done的定义
		// func (wg *WaitGroup) Done() {
		// 	wg.Add(-1)
		// }
		go func() {
			wg.Done()
			time.Sleep(time.Second * 1)
		}()
	}

	wg.Wait() //
	fmt.Println("wait group exit")
}
