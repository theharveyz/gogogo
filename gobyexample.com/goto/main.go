package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func trafficLight() {
RED:
	fmt.Println("The red light is on.")
	goto GREEN
YELLO:
	fmt.Println("The yello light is on.")
	goto RED
GREEN:
	fmt.Println("The green light is on.")
	goto YELLO

	// 只能在同一个代码块中使用, 下面的不被允许
	// go func() {
	// 	goto RED
	// }()
	goto RED
}

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	fmt.Println("starting...")
	go trafficLight()
	go func() {
		for {
			fmt.Println("hello...")
			time.Sleep(time.Second)
		}

	}()
	<-c
}
