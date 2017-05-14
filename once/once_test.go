package main

import "testing"
import "fmt"
import "sync"
import "time"

// sync.Once 内部通过读写锁计数器+控制函数只能调用一次

func TestOnce(t *testing.T) {
	var o sync.Once
	for i, v := range make([]string, 10) {
		o.Do(once)
		fmt.Println("range once...count:", v, "---", i)
	}
	for i := 0; i < 10; i++ {
		o.Do(onced)
		fmt.Println("onced...count:", i)
	}
	time.Sleep(4000)
}

// 调用结果, onced不会被调用
// sync.Once对象保证本身只会被调用一次, 即使里面的函数被替换(比如被onced替换, onced也不会执行)

func once() {
	fmt.Println("once...")
}

func onced() {
	fmt.Println("onced...")
}
