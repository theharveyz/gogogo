package main

import (
	"fmt"
	"time"
)

// 双向channel可以转换为单向channel, 作为函数参数要做严格限制
func worker(c chan<- bool) {
	fmt.Println("worker working")
	time.Sleep(time.Second * 3)
	c <- true
}

func ping(pings chan<- string, s string) {
	pings <- s
}

func pong(pings <-chan string, pongs chan<- string) {
	pongs <- (<-pings)
}

func main() {
	ch := make(chan string)
	go func() {
		ch <- "message"
	}()
	fmt.Println("chan: ", <-ch)

	buf_chan := make(chan string, 2)
	go func() {
		buf_chan <- "message"
	}()
	fmt.Println("chan: ", <-buf_chan)
	buf_chan <- "message"
	fmt.Println("chan: ", <-buf_chan)

	wc := make(chan bool, 1)
	go worker(wc)
	fmt.Println(<-wc)
	fmt.Println("worker done")

	pings, pongs := make(chan string, 1), make(chan string, 1)
	ping(pings, "ping...")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
