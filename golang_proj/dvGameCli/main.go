//客户端发送封包
package main

import (
	"fmt"
	"net"
	"os"
	"time"

	"protocol"
)

func sender(conn net.Conn) {
	for i := 0; i < 1000; i++ {
		words := "{\"Id\":1,\"Name\":\"golang\",\"Message\":\"message\"}"
		conn.Write(protocol.Packet([]byte(words)))
	}
	fmt.Println("send over")
}

func main() {
	server := "127.0.0.1:9989"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	defer conn.Close()
	fmt.Println("connect success")
	go sender(conn)
	for {
		time.Sleep(1 * 1e9)
	}
}
