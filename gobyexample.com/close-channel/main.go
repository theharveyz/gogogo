package main

import (
	"fmt"
)

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)
	go func() {
		for {

			if j, more := <-jobs; more {
				fmt.Println("job", j)
			} else {
				fmt.Println(more)
				fmt.Println("all jobs receiver")
				done <- true
			}
		}
	}()

	for i := 0; i < 3; i++ {
		fmt.Println("job:", i)
		jobs <- i
	}
	close(jobs)
	<-done
}
