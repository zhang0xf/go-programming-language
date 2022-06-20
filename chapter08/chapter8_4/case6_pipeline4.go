package chapter8_4

import "fmt"

// 带缓存的Channels

// 带缓存的Channel内部持有一个元素队列。队列的最大容量是在调用make函数创建channel时通过第二个参数指定的。
// 下面的语句创建了一个可以持有三个字符串元素的带缓存Channel。
func BufferedChannels() {
	ch := make(chan string, 3)
	close(ch)
}

// 向缓存Channel的发送操作就是向内部缓存队列的尾部插入元素，接收操作则是从队列的头部删除元素。
// 如果内部缓存队列是满的，那么发送操作将阻塞直到因另一个goroutine执行接收操作而释放了新的队列空间。
// 相反，如果channel是空的，接收操作将阻塞直到有另一个goroutine执行发送操作而向队列插入元素。
func BufferedChannels2() {
	ch := make(chan string, 3)

	// 我们可以在无阻塞的情况下连续向新创建的channel发送三个值：
	ch <- "A"
	ch <- "B"
	ch <- "C"

	// 此刻，channel的内部缓存队列将是满的（图8.3），如果有第四个发送操作将发生阻塞。
	// 如果我们接收一个值:
	fmt.Println(<-ch) // "A"
	// 那么channel的缓存队列将不是满的也不是空的（图8.4），因此对该channel执行的发送或接收操作都不会发生阻塞。
	// 通过这种方式，channel的缓存队列解耦了接收和发送的goroutine。(只管接和发,空或满会自动阻塞)

	// 在某些特殊情况下，程序可能需要知道channel内部缓存的容量，可以用内置的cap函数获取：
	fmt.Println(cap(ch)) // "3"

	// 同样，对于内置的len函数，如果传入的是channel，那么将返回channel内部缓存队列中有效元素的个数。
	fmt.Println(len(ch)) // "2"

	// 在继续执行两次接收操作后channel内部的缓存队列将又成为空的，如果有第四个接收操作将发生阻塞：
	fmt.Println(<-ch) // "B"
	fmt.Println(<-ch) // "C"

	close(ch)
}

// 在这个例子中，发送和接收操作都发生在同一个goroutine中，但是在真实的程序中它们一般由不同的goroutine执行。
// Go语言新手有时候会将一个带缓存的channel当作同一个goroutine中的队列使用，虽然语法看似简单，但实际上这是一个错误。
// Channel和goroutine的调度器机制是紧密相连的，如果没有其他goroutine从channel接收，发送者——或许是整个程序——将会面临永远阻塞的风险。
// 如果你只是需要一个简单的队列，使用slice就可以了。

// 下面的例子展示了一个使用了带缓存channel的应用。它并发地向三个镜像站点发出请求，三个镜像站点分散在不同的地理位置。
// 它们分别将收到的响应发送到带缓存channel，最后接收者只接收第一个收到的响应，也就是最快的那个响应。
// 因此mirroredQuery函数可能在另外两个响应慢的镜像站点响应之前就返回了结果。
// (顺便说一下，多个goroutines并发地向同一个channel发送数据，或从同一个channel接收数据都是常见的用法。)
func mirroredQuery() string {
	responses := make(chan string, 3)
	go func() { responses <- request("asia.gopl.io") }()
	go func() { responses <- request("europe.gopl.io") }()
	go func() { responses <- request("americas.gopl.io") }()
	return <-responses // return the quickest response
}

func request(hostname string) (response string) {
	/* ... */
	return ""
}

// 如果我们使用了无缓存的channel，那么两个慢的goroutines将会因为没有人接收而被永远卡住。这种情况，称为goroutines泄漏，这将是一个BUG。
// 和垃圾变量不同，泄漏的goroutines并不会被自动回收，因此确保每个不再需要的goroutine能正常退出是重要的。
