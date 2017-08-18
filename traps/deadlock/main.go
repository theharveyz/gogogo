package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var lock = &sync.Mutex{}

func init() {
	runtime.GOMAXPROCS(10)
}

func main() {
	lock.Lock()
	defer lock.Unlock()
	go func() {
		defer func() {
			if e := recover(); e != nil {
				fmt.Println(e)
			}
		}()
		lock.Lock()
		fmt.Println("test")
		panic("test")
		lock.Unlock()
	}()
	<-time.After(time.Second * 50)
	return
}
