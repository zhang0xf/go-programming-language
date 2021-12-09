package chapter8_2

import (
	"log"
	"net"
)

// 现在多个客户端可以同时接收到时间了

func Clock2() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		// 我们这里对服务端程序做一点小改动，使其支持并发：在handleConn函数调用的地方增加go关键字，让每一次handleConn的调用都进入一个独立的goroutine。
		go handleConn(conn) // handle one connection at a time
	}
}
