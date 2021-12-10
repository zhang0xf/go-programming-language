package chapter8_6

import (
	"fmt"
	"log"
	"os"

	"gopl.io/ch5/links"
)

// 为了解决无穷无尽的并行的问题，我们可以限制并发程序所使用的资源来使之适应自己的运行环境。
// 我们可以用一个有容量限制的buffered channel来控制并发，这类似于操作系统里的计数信号量概念。
// 由于channel里的元素类型并不重要，我们用一个零值的struct{}来作为其元素。
// 让我们重写crawl函数，将对links.Extract的调用操作用获取、释放token的操作包裹起来，来确保*同一时间*对其只有20个调用。信号量数量和其能操作的IO资源数量应保持接近。
// 为了使这个程序能够终止，我们需要在worklist为空或者没有crawl的goroutine在运行时退出主循环。

// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 20)

func crawl2(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release the token
	if err != nil {
		log.Print(err)
	}
	return list
}

// 这个版本中，计数器n对worklist的发送操作数量进行了限制。每一次我们发现有元素需要被发送到worklist时，我们都会对n进行++操作，在向worklist中发送初始的命令行参数之前，我们也进行过一次++操作。
// 这里的操作++是在每启动一个crawler的goroutine之前。主循环会在n减为0时终止，这时候说明没活可干了。
func Crawl2() {
	worklist := make(chan []string)
	var n int // number of pending sends to worklist

	// Start with the command-line arguments.
	n++
	go func() { worklist <- os.Args[1:] }()

	// Crawl the web concurrently.
	seen := make(map[string]bool)

	for ; n > 0; n-- {
		list := <-worklist // 为空：程序会阻塞,但是这里不可能为空！因为计数器n，没活可干时，不会进入for循环
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl2(link)
				}(link)
			}
		}
	}
}
