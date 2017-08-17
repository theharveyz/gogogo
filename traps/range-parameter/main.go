package main

import (
	"fmt"
	"sync"
)

func main() {
	ss := map[string]string{
		"a": "a",
		"b": "b",
		"c": "c",
	}
	wg := &sync.WaitGroup{}
	for key, v := range ss {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(key)
			fmt.Println(v)
			// c
			// c
			// c
			// c
			// c
			// c
		}()
	}
	wg.Wait()
}
