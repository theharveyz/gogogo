package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	c1, c2 := make(chan string, 1), make(chan string, 1)
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
			// default:
			// 	fmt.Println("select default")
		}
	}
	go func() {
		for {
			fmt.Println("empty select1")
			select {}
			fmt.Println("empty select2")
		}
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c
}
