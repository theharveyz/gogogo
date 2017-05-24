package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"
)

const (
	MAX_SIZE                  = 6 * 1024
	MIN_GRACEFUL_SECONDS      = 3
	LF                   rune = 10
	CR                   rune = 13
)

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

func (udps *UDPServer) GracefulShutdown(d time.Duration) {
	udps.mu.Lock()
	udps.shutdown = true
	// 立即释放锁
	udps.mu.Unlock()

	// 给定足够长时间,让未执行完的handler执行完
	if d < MIN_GRACEFUL_SECONDS*time.Second {
		d = MIN_GRACEFUL_SECONDS * time.Second
	}
	<-time.After(d)
}

func udpHandler(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return nil
	default:
	}
	data, is := ctx.Value("data").([]byte)
	if is {
		fmt.Println(strings.TrimFunc(string(data), func(b rune) bool {
			return (b == 0) || (b == LF) || (b == CR)
		}))
	}
	return nil
}

func main() {
	sign := make(chan os.Signal)
	exit := make(chan bool)
	signal.Notify(sign, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)

	udps := NewUDPServer()

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	go func() {
		<-sign
		exit <- true
	}()
	go func() {
		// 这里要求是阻塞的
		udps.Listen(9981)
		exit <- true
	}()

	<-exit
	udps.GracefulShutdown(time.Second * 3)
	now := time.Now()
	fmt.Printf("\nThe udp server shutdown at: %d-%d-%d %d:%d:%d", now.Year(),
		now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	// todo
}
