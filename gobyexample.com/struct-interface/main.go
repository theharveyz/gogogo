package main

import (
	"fmt"

	"gogogo/gobyexample.com/struct-interface/t"
)

type Bar struct {
	test.T
}

func main() {
	fmt.Println(&Bar{})
}
