package main

// 不能关闭只出不进的channel
func main() {
	ch := make(chan int, 1)
	go func(c <-chan int) {
		close(c) // invalid operation: close(c) (cannot close receive-only channel)
	}(ch)

	<-ch
}
