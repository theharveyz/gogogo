package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// static file server
	http.Handle("/", http.FileServer(http.Dir("./static")))
	go func() {
		if err := http.ListenAndServe(":1901", nil); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}()
	fmt.Println("http server start in port:1901")

	nc := make(chan os.Signal, 1)
	signal.Notify(nc, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, os.Interrupt)
	select {
	case s := <-nc:
		fmt.Println("signal recieve:", s)
		os.Exit(1)
	}
}
