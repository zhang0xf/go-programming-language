package chapter8_10

// 内部变量clients会记录当前建立连接的客户端集合。其记录的内容是每一个客户端的消息发出channel的"资格"信息。
// broadcaster监听来自全局的entering和leaving的channel来获知客户端的到来和离开事件。
// broadcaster也会监听全局的消息channel，所有的客户端都会向这个channel中发送消息。

type client chan<- string // an outgoing message channel

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}
