package main

import (
	"fmt"
	"time"
)

func main() {

	// timer解决一次性需求, 类似于setTimeout
	timer1 := time.NewTimer(time.Second * 2)

	<-timer1.C
	fmt.Println("timer1 expired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("timer2 expired")
	}()

	// 终止timer
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("timer2 stopped")
	}

	time.Since(time.Now())
	// 未能及时终止
	timer3 := time.NewTimer(time.Second)
	time.Sleep(time.Second * 2)
	stop3 := timer3.Stop()
	fmt.Println(stop3) // false
}
