package main

import (
	"fmt"
	"os"
)

func main() {
	// 通过os.Exit退出时, defer将不会执行
	defer fmt.Println("!")

	// go语言不像C一样将main函数返回值表示退出状态, go语言需要使用os.Exit()才行
	os.Exit(3)

	// echo $?
	// output: 3
}
