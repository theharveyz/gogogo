package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		// ---- 正确写法:
		// go func(i int) {
		// 	defer wg.Done()
		// 	fmt.Println(i)
		// }(i)

		// 共享内存i ----- 要注意!!!
		go func() {
			defer wg.Done()
			fmt.Println(i)
		}()
		// 10
		// 10
		// 10
		// 10
		// ...
	}
	wg.Wait()
}
