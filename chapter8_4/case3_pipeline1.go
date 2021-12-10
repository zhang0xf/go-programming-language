package chapter8_4

import (
	"fmt"
)

// Channels也可以用于将多个goroutine连接在一起，一个Channel的输出作为下一个Channel的输入。这种串联的Channels就是所谓的管道（pipeline）。

// 第一个goroutine是一个计数器，用于生成0、1、2、……形式的整数序列，然后通过channel将该整数序列发送给第二个goroutine；
// 第二个goroutine是一个求平方的程序，对收到的每个整数求平方，然后将平方后的结果通过第二个channel发送给第三个goroutine；
// 第三个goroutine是一个打印程序，打印收到的每个整数。

func Pipeline1() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; ; x++ {
			naturals <- x
			fmt.Printf("naturals被写入\n")
		}
	}()

	// Squarer
	go func() {
		for {
			x := <-naturals
			fmt.Printf("从naturals读取\n")
			squares <- x * x
			fmt.Printf("squares被写入\n")
		}
	}()

	// Printer (in main goroutine)
	for {
		fmt.Printf("从squares读取\n")
		fmt.Println(<-squares)
	}
}
