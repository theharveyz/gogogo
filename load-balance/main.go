package main

import (
	"fmt"

	"./roundrobin" // import pkg
)

func main() {
	rr := &round_robin.RoundRobin{}
	rr.AddNode(1, 2, 3, 4)

	fmt.Println(rr.GetNode())
	fmt.Println(rr.GetNode())
	fmt.Println(rr.GetNode())
	fmt.Println(rr.GetNode())
	fmt.Println(rr.GetNode())
	fmt.Println(rr.GetNode())
	rr.AddNode(5)
	rr.RemoveNode(2)
	fmt.Println(rr.GetNode())
	fmt.Println(rr.GetNode())
	fmt.Println(rr.GetNode())
	fmt.Println(rr.GetNode())
	fmt.Println(rr.GetNode())
	fmt.Println(rr.GetNode())
	fmt.Println(rr.SeekIndex())
	fmt.Println(rr.SeekIndex())
}
