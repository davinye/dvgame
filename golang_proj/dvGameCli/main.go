package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

var nick string = ""

func main() {
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:9989")
	checkErr(err)
	conn, err := net.DialTCP("tcp", nil, addr)
	checkErr(err)

	println(conn.RemoteAddr().String())

	// 读取提示
	data := make([]byte, 50)
	conn.Read(data)
	fmt.Println(string(data))

	conn.Write([]byte("我来了"))
	go Handle(conn)

	/*	// 输入昵称
		fmt.Print("输入昵称:")
		fmt.Scanf("%v", &nick)
		fmt.Println("Hello " + nick)
		conn.Write([]byte("nick|" + nick))

		go Handle(conn)

		for {
			someTex := ""
			fmt.Scanf("%v", &someTex)
			conn.Write([]byte("say|" + nick + "|" + someTex))
		}*/
}

const BufLength = 128

func Handle(conn net.Conn) {
	headBuff := make([]byte, 50) // set read stream size

	conn.Read(headBuff)
	println(string(headBuff))

	return

	for {
		data := make([]byte, 1024)
		buf := make([]byte, BufLength)
		for {
			n, err := conn.Read(buf)
			if err != nil && err != io.EOF {
				checkErr(err)
			}
			data = append(data, buf[:n]...)
			if n != BufLength {
				break
			}
		}

		fmt.Println(string(data))
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
