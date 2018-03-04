package main

import (
	"fmt"
)

func main() {
	// rune 能支持任何字符, 而byte不支持中文
	fmt.Println(len([]byte("你好"))) //output: 6
	fmt.Println(len([]rune("你好"))) //output: 2
}
