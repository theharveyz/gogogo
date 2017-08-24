package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for range time.Tick(time.Second) {
			fmt.Println("tick")
		}
	}()
	// for {
	// }
	// 等价于
	select {}
}
