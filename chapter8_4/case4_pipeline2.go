package chapter8_4

import "fmt"

// 如果发送者知道，没有更多的值需要发送到channel的话，那么让接收者也能及时知道没有多余的值可接收将是有用的，因为接收者可以停止不必要的接收等待。
// 当一个channel被关闭后，再向该channel发送数据将导致panic异常。
// 当一个被关闭的channel中已经发送的数据都被成功接收后，后续的接收操作将不再阻塞，它们会立即返回一个零值。
// 关闭pipeline1例子中的naturals变量对应的channel并不能终止循环，它依然会收到一个永无休止的零值序列，然后将它们发送给打印者goroutine。

// 没有办法直接测试一个channel是否被关闭，但是接收操作有一个变体形式：
// 它多接收一个结果，多接收的第二个结果是一个布尔值ok，ture表示成功从channels接收到值，false表示channels已经被关闭并且里面没有值可接收。
// 使用这个特性，我们可以修改squarer函数中的循环代码，当naturals对应的channel被关闭并没有值可接收时跳出循环，并且也关闭squares对应的channel.
func Pipeline1Variants() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; ; x++ {
			naturals <- x
		}
	}()

	// Squarer
	go func() {
		for {
			x, ok := <-naturals
			if !ok {
				break // channel was closed and drained
			}
			squares <- x * x
		}
		close(squares)
	}()

	// Printer (in main goroutine)
	for {
		fmt.Println(<-squares)
	}
}

// 因为上面的语法是笨拙的，而且这种处理模式很常见，因此Go语言的range循环可直接在channels上面迭代。
// 使用range循环是上面处理模式的简洁语法，它依次从channel接收数据，当channel被关闭并且没有值可接收时跳出循环。

// 在下面的改进中，我们的计数器goroutine只生成100个含数字的序列，然后关闭naturals对应的channel，这将导致计算平方数的squarer对应的goroutine可以正常终止循环并关闭squares对应的channel。
func Pipeline2() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; x < 100; x++ {
			fmt.Printf("写入naturals\n")
			naturals <- x
		}
		fmt.Printf("关闭naturals\n")
		close(naturals)
	}()

	// Squarer
	go func() {
		for x := range naturals {
			fmt.Printf("写入squares\n")
			squares <- x * x
		}
		close(squares)
	}()

	// Printer (in main goroutine)
	for x := range squares {
		fmt.Println(x)
	}
}

// 其实你并不需要关闭每一个channel。只有当需要告诉接收者goroutine，所有的数据已经全部发送时才需要关闭channel。
// 不管一个channel是否被关闭，当它没有被引用时将会被Go语言的垃圾自动回收器回收。(不要将关闭一个打开文件的操作和关闭一个channel操作混淆。文件属于系统资源,需要主动close)
// 试图重复关闭一个channel将导致panic异常，试图关闭一个nil值的channel也将导致panic异常。
// 关闭一个channels还会触发一个广播机制，我们将在8.9节讨论。
