package chapter8_10

import (
	"log"
	"net"
)

// usage : 需要多个进程
// ./exercise 启动服务器
// 使用netcat做客户端来聊天

// 我们用一个聊天服务器来终结本章节的内容，这个程序可以让一些用户通过服务器向其它所有用户广播文本消息。

// 当与n个客户端保持聊天session时，这个程序会有2n+2个并发的goroutine，然而这个程序却并不需要显式的锁(§9.2)。
// clients这个map被限制在了一个独立的goroutine中，broadcaster，所以它不能被并发地访问。
// 多个goroutine共享的变量只有这些channel和net.Conn的实例，两个东西都是并发安全的。
// 我们会在下一章中更多地讲解约束，并发安全以及goroutine中共享变量的含义。

func Chat() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
