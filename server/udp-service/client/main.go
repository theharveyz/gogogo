package main

import (
	"fmt"
	"net"
)

func main() {
	ip := net.ParseIP("127.0.0.1")
	fmt.Println(ip)
	fmt.Println(net.IPv4zero)

	srcAddr := &net.UDPAddr{IP: net.IPv4zero, Port: 0}
	dstAddr := &net.UDPAddr{IP: ip, Port: 9981}

	conn, err := net.DialUDP("udp", srcAddr, dstAddr)

	if err != nil {
		fmt.Println(err)
	}

	defer conn.Close()
	conn.Write([]byte(`
	{"a":"b"}
	`))
	fmt.Printf("<%s>", conn.RemoteAddr())
	fmt.Println([]byte("\r")) // 13
	fmt.Println([]byte("\n")) // 10
}
