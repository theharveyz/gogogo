package main

import (
	"fmt"
)

func fact(n int) int {
	if 0 == n {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))
}
