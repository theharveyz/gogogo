package main

import "os"
import "fmt"

func main() {
	exit := make(chan struct{})
	sigc := make(chan os.Signal, 1)

	select {
	case s := <-sigc:
		fmt.Println("receiver signal", s)
	case <-exit:
		fmt.Println("done...")
	}
}
