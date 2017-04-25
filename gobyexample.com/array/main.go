package main

import "fmt"

func pl(strs ...interface{}) {
	fmt.Println(strs...) // 变参
}

func main() {
	// empty array
	var em [5]int
	fmt.Println(em)

	oneD := [5]int{1, 2, 3, 4, 5}
	fmt.Println(oneD)
	fmt.Println(len(oneD))

	var twoD [2][3]int
	for i := 0; i < len(twoD); i++ {
		for n := 0; n < len(twoD[i]); n++ {
			twoD[i][n] = i + n
		}
	}
	pl(twoD)
}
