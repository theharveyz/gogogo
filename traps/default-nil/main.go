package main

import (
	"fmt"
)

type F func()

func main() {
	var f F
	fmt.Println(f) // default: nil

	var m map[int]int
	fmt.Println(m) // default: map[]

	var c chan int
	fmt.Println(c) // default: nil
}
