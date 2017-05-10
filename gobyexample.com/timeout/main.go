package main

import (
	"fmt"
	"time"
)

func main() {
	c1, c2 := make(chan string, 1), make(chan string, 1)

	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "hello"
	}()
	go func() {
		time.Sleep(time.Second * 3)
		c2 <- "go"
	}()
	select {
	case msg1 := <-c1:
		fmt.Println(msg1)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout")

	}

	select {
	case msg2 := <-c2:
		fmt.Println(msg2)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout")

	}
}
