package main

import (
	"os"
)

func main() {
	panic("a problem")

	// 错误
	_, err := os.Create("/tmp/file")
	// 错误触发的异常
	if err != nil {
		panic(err)
	}
}
