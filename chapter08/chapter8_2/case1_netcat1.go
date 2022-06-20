package chapter8_2

import (
	"io"
	"log"
	"net"
	"os"
)

// usage : 需要两个可执行程序,当前环境可能并不支持,但用法也显而易见.

// 也可以用我们下面的这个用go写的简单的telnet程序，用net.Dial就可以简单地创建一个TCP连接：
// 让我们同时运行两个客户端来进行一个测试，第二个客户端必须等待第一个客户端完成工作，这样服务端才能继续向后执行；因为我们这里的服务器程序clock1同一时间只能处理一个客户端连接。

func Netcat1() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
