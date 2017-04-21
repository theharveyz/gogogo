package main

import (
	"fmt"
)

func main() {
	i := 1
	for i < 3 {
		fmt.Println(i)
		i++
	}

	for ; i < 9; i++ {
		fmt.Println(i)
	}

	for j := 1; j < i; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println(i)
		break
	}

	var s, retry int = 1, 0
RETRY:
	for ; s < 3; s++ {
		if s%2 == 0 {
			fmt.Println(s)
		}
	}
	retry++
	if retry < s {
		fmt.Println(s)
		fmt.Println("retry ==>", retry)
		goto RETRY
	}

}
