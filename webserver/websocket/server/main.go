package main

import "runtime"

// go websocket server

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU() * 5)
}

func main() {

}
