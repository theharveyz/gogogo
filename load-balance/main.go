package main

import (
	"fmt"

	"./roundrobin" // import pkg
)

func main() {
	rr := &round_robin.RoundRobin{Index: 1, Keys: make(map[int]interface{})}
	fmt.Println(rr)
}
