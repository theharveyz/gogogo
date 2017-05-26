package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	a, b := "hello", "world"

	// 1
	fmt.Printf("%s%s", a, b)

	// 2
	fmt.Println(a + b)

	// fastest
	fmt.Println(strings.Join([]string{a, b}, ""))

	// byte buffer
	var buffer bytes.Buffer
	buffer.WriteString(a)
	buffer.WriteString(b)
	fmt.Println(buffer.String())
}
