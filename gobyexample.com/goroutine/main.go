package main

import (
	"fmt"
)

func f(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s, ":", i)
	}
}

func main() {
	f("start")
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	go f("goroutine")

	var intput string
	fmt.Scanln(&intput)
	fmt.Println("done")
}
