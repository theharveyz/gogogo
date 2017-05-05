package main

import (
	"fmt"
)

func main() {
	var i, s int = 1, 2
	var demos [1]string
	fmt.Println(i + s)
	fmt.Println(&demos[0])

	var e int
	var str string
	var n map[string]int

	fmt.Println(e, str, " --") // 0   --
	fmt.Println(n)
}
