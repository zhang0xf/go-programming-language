package chapter8_7

import (
	"fmt"
	"time"
)

// 基于select的多路复用

// 下面的程序会进行火箭发射的倒计时。
// time.Tick函数返回一个channel，程序会周期性地像一个节拍器一样向这个channel发送事件。
// 每一个事件的值是一个时间戳，不过更有意思的是其传送方式。

func Countdown1() {
	fmt.Println("Commencing countdown.")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		<-tick
	}
	launch()
}

func launch() {
	// ...
}
