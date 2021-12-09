// Clock1 is a TCP server that periodically writes the time.
package chapter8_2

import (
	"io"
	"log"
	"net"
	"time"
)

// usage : ./exercise (start the server)
//         nc localhost 8000 (client)

// 网络编程是并发大显身手的一个领域，由于服务器是最典型的需要同时处理很多连接的程序，这些连接一般来自于彼此独立的客户端。
// 我们的第一个例子是一个顺序执行的时钟服务器，它会每隔一秒钟将当前时间写到客户端：
// 为了连接例子里的服务器，我们需要一个客户端程序，比如netcat这个工具(nc命令)，这个工具可以用来执行网络连接操作。
// 如果你的系统没有装nc这个工具，你可以用telnet来实现同样的效果

func Clock1() {
	// Listen函数创建了一个net.Listener的对象，这个对象会监听一个网络端口上到来的连接
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
		handleConn(conn) // handle one connection at a time
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		// time.Time.Format方法提供了一种格式化日期和时间信息的方式。
		// 它的参数是一个格式化模板，标识如何来格式化时间，而这个格式化模板限定为Mon Jan 2 03:04:05PM 2006 UTC-0700。
		// 在例子中我们只用到了小时、分钟和秒。
		// time包里定义了很多标准时间格式，比如time.RFC1123。
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}
