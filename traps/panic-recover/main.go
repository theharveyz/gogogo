package main

import (
	"errors"
	"fmt"
)

func main() {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()

		go func() {
			panic(errors.New("oh fuck panic!")) // 该panic不会被捕获!!!!
		}()
	}()

	select {}
}
