package chapter8_7

import (
	"fmt"
	"os"
	"time"
)

// 现在我们让这个程序支持在倒计时中，用户按下return键时直接中断发射流程。
// 首先，我们启动一个goroutine，这个goroutine会尝试从标准输入中读入一个单独的byte并且，如果成功了，会向名为abort的channel发送一个值。

func Countdown2() {
	fmt.Println("Commencing countdown.")

	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		abort <- struct{}{}
	}()

	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		<-tick
	}

	launch()
}

// 我们无法做到从每一个channel中接收信息，如果我们这么做的话，如果第一个channel中没有事件发过来那么程序就会立刻被阻塞，这样我们就无法收到第二个channel中发过来的事件。
// 这时候我们需要多路复用(multiplex)这些操作了，为了能够多路复用，我们使用了select语句。
// 每一个case代表一个通信操作（在某个channel上进行发送或者接收），并且会包含一些语句组成的一个语句块。
// 一个接收表达式可能只包含接收表达式自身，就像下面的第一个case，或者包含在一个简短的变量声明中，像第二个case里一样；第二种形式让你能够引用接收到的值。
func SelectSample() {
	ch1 := make(chan int)
	ch2 := make(chan []string)
	ch3 := make(chan time.Time)

	y := time.Now()

	select {
	case <-ch1:
		// ...
	case x := <-ch2:
		// ...use x...
		fmt.Printf("%T\n", x)
	case ch3 <- y:
		// ...
	default:
		// ...
	}
}
