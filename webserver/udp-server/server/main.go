package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func init() {

}

func parse(ctx context.Context, data []byte) {
	select {
	case <-ctx.Done():
		return
	}

}

var lock = new(sync.RWMutex)

func udpSrv(ctx context.Context, exit chan<- bool) {

	listener, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 9981})
	if err != nil {
		fmt.Println(err)
		exit <- true
	}

	shutdown := false

	for {
		lock.RLock()
		if shutdown {
			break
		}
		lock.RUnlock()
		data := make([]byte, 1024) // 重新赋值
		n, remoteAddr, err := listener.ReadFromUDP(data)
		fmt.Println("remote:", remoteAddr)
		go parse(ctx, data)
	}

	defer func() {
		// close
	}()

	return listener, func() {
		defer lock.Unlock()
		lock.Lock()
		shutdown = true
	}
}

func main() {
	sign := make(chan os.Signal)
	exit := make(chan bool)
	signal.Notify(sign, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		<-sign
		exit <- true
	}()
	go udpSrv(ctx, exit)

	<-exit
	defer cancel()
	// todo
}
