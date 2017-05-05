package main

import (
	"fmt"

	"./roundrobin" // import pkg
)

func main() {
	rr := &round_robin.RoundRobin{1, make(map[int]interface{}, 0)}
	fmt.Println(&round_robin.RoundRobin)
}
