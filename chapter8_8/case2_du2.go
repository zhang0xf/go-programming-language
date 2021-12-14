package chapter8_8

import (
	"flag"
	"time"
)

// usage : ./exercise $HOME /usr /bin /etc

// 下面这个du的变种会间歇打印内容，不过只有在调用时提供了-v的flag才会显示程序进度信息。
// 在roots目录上循环的后台goroutine在这里保持不变。
// 主goroutine现在使用了计时器来每500ms生成事件，然后用select语句来等待文件大小的消息来更新总大小数据，或者一个计时器的事件来打印当前的总大小数据。
// 如果-v的flag在运行时没有传入的话，tick这个channel会保持为nil，这样在select里的case也就相当于被禁用了。
// 由于我们的程序不再使用range循环，第一个select的case必须显式地判断fileSizes的channel是不是已经被关闭了，这里可以用到channel接收的二值形式。如果channel已经被关闭了的话，程序会直接退出循环。
// 这里的break语句用到了标签break，这样可以同时终结select和for两个循环；如果没有用标签就break的话只会退出内层的select循环，而外层的for循环会使之进入下一轮select循环。
// 然而这个程序还是会花上很长时间才会结束。完全可以并发调用walkDir，从而发挥磁盘系统的并行性能。

var verbose = flag.Bool("v", false, "show verbose progress messages")

func Du2() {
	// Determine the initial directories.
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// ...start background goroutine...
	fileSizes := make(chan int64)
	go func() {
		for _, root := range roots {
			walkDir(root, fileSizes)
		}
		close(fileSizes)
	}()

	// Print the results periodically.
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	var nfiles, nbytes int64
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes was closed
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}
	printDiskUsage(nfiles, nbytes) // final totals
}
