package main

import (
	"fmt"
)

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
}
