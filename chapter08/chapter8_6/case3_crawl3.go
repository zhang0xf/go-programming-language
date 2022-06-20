package chapter8_6

import (
	"fmt"
	"os"
)

// 下面的程序是避免过度并发的另一种思路。
// 这个版本使用了原来的crawl函数，但没有使用计数信号量，取而代之用了20个常驻的crawler goroutine，这样来保证最多20个HTTP请求在并发。
// 为了节省篇幅，这个例子的终止问题我们先不进行详细阐述了。

func Crawl3() {
	worklist := make(chan []string)  // lists of URLs, may have duplicates
	unseenLinks := make(chan string) // de-duplicated URLs

	// Add command-line arguments to worklist.
	go func() { worklist <- os.Args[1:] }()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() { // go关键字使得匿名函数在unseenLinks为空时，不会阻塞for循环。
			for link := range unseenLinks {
				foundLinks := crawl1(link)
				// go func() { worklist <- foundLinks }() // crawl函数爬到的链接在一个专有的goroutine中被发送到worklist中来避免死锁。?
				fmt.Printf("写worklist\n")
				worklist <- foundLinks // 死锁:主线程往unseenLinks写，20个子线程往worklist写；当20个子线程都卡在写worklist处，主线程就会卡在写unseenLinks处（此时，unseenLinks没有了读）
				fmt.Printf("写完成worklist...\n")
			}
		}()
	}

	// 主goroutine负责拆分它从worklist里拿到的元素，然后把没有抓过的经由unseenLinks channel发送给一个爬虫的goroutine。
	// seen这个map被限定在main goroutine中；也就是说这个map只能在main goroutine中进行访问。
	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for list := range worklist {
		fmt.Printf("读worklist\n")
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link // 死锁:主线程往unseenLinks写，20个子线程往worklist写；当20个子线程都卡在写worklist处，主线程就会卡在写unseenLinks处（此时，unseenLinks没有了读）
			}
		}
		fmt.Printf("读完成worklist...\n")
	}
}
