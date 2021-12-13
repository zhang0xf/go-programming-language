package chapter8_7

import "fmt"

// 有时候我们希望能够从channel中发送或者接收值，并避免因为发送或者接收导致的阻塞，尤其是当channel没有准备好写或者读时。
// select语句就可以实现这样的功能。select会有一个default来设置当其它的操作都不能够马上被处理时程序需要执行哪些逻辑。

// 下面的select语句会在abort channel中有值时，从其中接收值；无值时什么都不做。
// 这是一个非阻塞的接收操作；反复地做这样的操作叫做“轮询channel”。

func SelectChannel2() {
	abort := make(chan struct{})

	select {
	case <-abort:
		fmt.Printf("Launch aborted!\n")
		return
	default:
		// do nothing
	}
}

// channel的零值是nil。也许会让你觉得比较奇怪，nil的channel有时候也是有一些用处的。
// 因为对一个nil的channel发送和接收操作会永远阻塞，在select语句中操作nil的channel永远都不会被select到。
// 这使得我们可以用nil来激活或者禁用case，来达成处理其它输入或输出事件时超时和取消的逻辑。我们会在下一节中看到一个例子。
