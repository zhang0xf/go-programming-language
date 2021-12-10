package chapter8_6

import (
	"fmt"
	"log"
	"os"

	"gopl.io/ch5/links"
)

//  usage : ./exercise http://www.baidu.com

// 在5.6节中，我们做了一个简单的web爬虫，用bfs(广度优先)算法来抓取整个网站。
// 在本节中，我们会让这个爬虫并行化，这样每一个彼此独立的抓取命令可以并行进行IO，最大化利用网络资源。
// crawl函数和gopl.io/ch5/findlinks3中的是一样的。

func crawl1(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

// 主函数和5.6节中的breadthFirst(广度优先)类似。
// 像之前一样，一个worklist是一个记录了需要处理的元素的队列，每一个元素都是一个需要抓取的URL列表，不过这一次我们用channel代替slice来做这个队列。
// 每一个对crawl的调用都会在他们自己的goroutine中进行并且会把他们抓到的链接发送回worklist。
// 另外注意这里将命令行参数传入worklist也是在一个另外的goroutine中进行的，这是为了避免channel两端的main goroutine与crawler goroutine都尝试向对方发送内容，却没有一端接收内容时发生死锁。
func Crawl1() {
	worklist := make(chan []string)

	// Start with the command-line arguments.
	go func() { worklist <- os.Args[1:] }()
	// worklist <- os.Args[1:] // 死锁:需要有其他线程读取worklist。

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				// 注意这里的crawl所在的goroutine会将link作为一个显式的参数传入，来避免“循环变量快照”的问题(在5.6.1中有讲解)。
				go func(link string) {
					worklist <- crawl1(link)
				}(link)
			}
		}
	}
}

// 无穷无尽地并行化并不是什么好事情，因为不管怎么说，你的系统总是会有一些个限制因素(会报错)
// 第二个问题是这个程序永远都不会终止，即使它已经爬到了所有初始链接衍生出的链接。
