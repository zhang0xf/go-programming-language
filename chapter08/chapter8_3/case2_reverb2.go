package chapter8_3

import (
	"bufio"
	"log"
	"net"
	"time"
)

// echo server -> netcat.go client

func Reverb2() {
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
		go handleConn2(conn) // handle one connection at a time
	}
}

// 客户端的每一次输入,服务器都开一个线程,使得客户端模拟"回声"更加真实!否则会出现这样的情形:客户端的第三次shout在前一个shout处理完成之前一直没有被处理.
// go后跟的函数的参数会在go语句自身执行时被求值；因此input.Text()会在main goroutine中被求值。
func handleConn2(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		go echo(c, input.Text(), 1*time.Second)
	}
	// NOTE: ignoring potential errors from input.Err()
	c.Close()
}

// 让服务使用并发不只是处理多个客户端的请求，甚至在处理单个连接时也可能会用到，就像我们上面的两个go关键词的用法。
// 然而在我们使用go关键词的同时，需要慎重地考虑net.Conn中的方法在并发地调用时是否安全，事实上对于大多数类型来说也确实不安全。我们会在下一章中详细地探讨并发安全性。
