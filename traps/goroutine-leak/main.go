package main

import (
	"flag"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"time"
)

var (
	PORT string
)

func init() {
	PORT = *flag.String("port", "6060", "pprof port")
}
func prof() {
	fmt.Println("pprof http port:", PORT)
	fmt.Println(http.ListenAndServe("localhost:"+string(PORT), nil))
}

func main() {
	ch := make(chan int, 1)
	go func() {
		for i := 0; i < 10; i++ {
			go func() {
				<-ch
			}()
		}
	}()

	// goroutine leak 资源泄露:
	go prof()
	// for循环
	for {
		time.Sleep(time.Second)
		runtime.GC() // 强制垃圾回收
	}
}
