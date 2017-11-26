package main

import (
	"fmt"
)

// 多个init函数则分别执行

func init() {
	fmt.Println("init 1")
}

func init() {
	fmt.Println("init 2")
}

func main() {
	fmt.Println("done")
}

// init 1
// init 2
// done
