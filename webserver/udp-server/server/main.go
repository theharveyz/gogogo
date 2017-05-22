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

const MAX_SIZE = 6 * 1024

type (
	UDPServer struct {
		shutdown bool
		listener *net.UDPConn
		mu       *sync.RWMutex
		ctx      context.Context
		cancel   context.CancelFunc
	}
)

func NewUDPServer() *UDPServer {
	ctx, cancel := context.WithCancel(context.Background())
	return &UDPServer{
		mu:     new(sync.RWMutex),
		ctx:    ctx,
		cancel: cancel,
	}
}

func (udps *UDPServer) Listen(port int) {
	listener, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: port})
	// 必须启动才可以
	if err != nil {
		panic(err)
	}
	udps.listener = listener

	fmt.Println("udp server listen successful at port ", port)

	for {
		// 这里不需要加锁
		if udps.IsShutdown() {
			fmt.Println("udp server shutdown...")
			break
		}
		data := make([]byte, MAX_SIZE) // 重新赋值
		if n, remoteAddr, err := udps.listener.ReadFromUDP(data); err != nil {
			fmt.Println("remote:", err)
		} else {
			fmt.Println("remote:", remoteAddr)
			fmt.Println("n:", n)
			newCtx := context.WithValue(udps.ctx, "data", data)
			go udpHandler(newCtx)
		}
	}
}

func (udps *UDPServer) IsShutdown() bool {
	defer udps.mu.RUnlock()
	udps.mu.RLock()
	return udps.shutdown
}

func (udps *UDPServer) Shutdown() {
	defer udps.mu.Unlock()
	udps.mu.Lock()
	udps.shutdown = true
}

func (udps *UDPServer) GracefulShutdown() {

}

func udpHandler(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return nil
	}
	fmt.Println(ctx)
	return nil
}

func main() {
	sign := make(chan os.Signal)
	exit := make(chan bool)
	signal.Notify(sign, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)

	udps := NewUDPServer()

	go func() {
		<-sign
		exit <- true
	}()
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
		// 这里要求是阻塞的
		udps.Listen(1992)
		exit <- true
	}()

	<-exit
	udps.Shutdown()
	// todo
}
