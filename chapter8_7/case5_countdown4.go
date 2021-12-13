package chapter8_7

import (
	"fmt"
	"time"
)

// 下面让我们的发射程序打印倒计时。这里的select语句会使每次循环迭代等待一秒来执行退出操作。

func CountDown4() {
	// ...create abort channel...
	abort := make(chan struct{})

	fmt.Println("Commencing countdown.  Press return to abort.")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick:
			// Do nothing.
		case <-abort:
			fmt.Println("Launch aborted!")
			return
		}
	}
	launch()
}

// time.Tick函数表现得好像它创建了一个在循环中调用time.Sleep的goroutine，每次被唤醒时发送一个事件。
// 当countdown函数返回时，它会停止从tick中接收事件，但是ticker这个goroutine还依然存活，继续徒劳地尝试向channel中发送值，然而这时候已经没有其它的goroutine会从该channel中接收值了--这被称为goroutine泄露(§8.4.4)。
// Tick函数挺方便，但是只有当程序整个生命周期都需要这个时间时我们使用它才比较合适。否则的话，我们应该使用下面的这种模式：

func TickStop() {
	ticker := time.NewTicker(1 * time.Second)
	<-ticker.C    // receive from the ticker's channel
	ticker.Stop() // cause the ticker's goroutine to terminate
}
