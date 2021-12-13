package chapter8_7

import (
	"fmt"
	"os"
	"time"
)

// select会等待case中有能够执行的case时去执行。
// 当条件满足时，select才会去通信并执行case之后的语句；这时候其它通信是不会执行的。
// 一个没有任何case的select语句写作select{}，会永远地等待下去。

// time.After函数会立即返回一个channel，并起一个新的goroutine在经过特定的时间后向该channel发送一个独立的值。
// 下面的select语句会一直等待直到两个事件中的一个到达，无论是abort事件或者一个10秒经过的事件。
// 如果10秒经过了还没有abort事件进入，那么火箭就会发射。

func CountDown3() {
	// ...create abort channel...
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		abort <- struct{}{}
	}()

	fmt.Println("Commencing countdown.  Press return to abort.")
	select {
	case <-time.After(10 * time.Second):
		// Do nothing.
	case <-abort:
		fmt.Println("Launch aborted!")
		return
	}
	launch()
}
