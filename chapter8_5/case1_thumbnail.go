package chapter8_5

import (
	"log"
	"os"
	"sync"

	"gopl.io/ch8/thumbnail"
)

// gopl.io/ch8/thumbnail包提供了ImageFile函数来帮我们拉伸图片。
// 显然我们处理文件的顺序无关紧要，因为每一个图片的拉伸操作和其它图片的处理操作都是彼此独立的。
// 像这种子问题都是完全彼此独立的问题被叫做易并行问题。易并行问题是最容易被实现成并行的一类问题，并且最能够享受到并发带来的好处，能够随着并行的规模线性地扩展。
// makeThumbnails makes thumbnails of the specified files.
func makeThumbnails1(filenames []string) {
	for _, f := range filenames {
		if _, err := thumbnail.ImageFile(f); err != nil {
			log.Println(err)
		}
	}
}

// 下面让我们并行地执行这些操作，从而将文件IO的延迟隐藏掉，并用上多核cpu的计算能力来拉伸图像。
// 即使当文件名的slice中只包含有一个元素。也比最早的版本使用的时间要短得多.为何?
// 答案其实是makeThumbnails在它还没有完成工作之前就已经返回了。它启动了所有的goroutine，每一个文件名对应一个，但没有等待它们一直到执行完毕。
// NOTE: incorrect!
func makeThumbnails2(filenames []string) {
	for _, f := range filenames {
		go thumbnail.ImageFile(f) // NOTE: ignoring errors
	}
}

// 没有什么直接的办法能够等待goroutine完成，但是我们可以改变goroutine里的代码让其能够将完成情况报告给外部的goroutine知晓，使用的方式是向一个共享的channel中发送事件。
// 因为我们已经确切地知道有len(filenames)个内部goroutine，所以外部的goroutine只需要在返回之前对这些事件计数。
// makeThumbnails3 makes thumbnails of the specified files in parallel.
func makeThumbnails3(filenames []string) {
	ch := make(chan struct{})
	for _, f := range filenames {
		// 注意我们将f的值作为一个显式的变量传给了函数，而不是在循环的闭包中声明：
		// 回忆一下之前在5.6.1节中，匿名函数中的循环变量快照问题。
		// 显式地添加这个参数，我们能够确保使用的f是当go语句执行时的“当前”那个f。
		go func(f string) {
			thumbnail.ImageFile(f) // NOTE: ignoring errors
			ch <- struct{}{}
		}(f)
	}

	// Wait for goroutines to complete.
	for range filenames {
		<-ch
	}
}

// 如果我们想要从每一个worker goroutine往主goroutine中返回值时该怎么办呢？
// 这个程序有一个微妙的bug。当它遇到第一个非nil的error时会直接将error返回到调用方，使得没有一个goroutine去排空errors channel。
// 这样剩下的worker goroutine在向这个channel中发送值时，都会永远地阻塞下去，并且永远都不会退出。
// 这种情况叫做goroutine泄露(§8.4.4)，可能会导致整个程序卡住或者跑出out of memory的错误。
// makeThumbnails4 makes thumbnails for the specified files in parallel.
// It returns an error if any step failed.
func makeThumbnails4(filenames []string) error {
	errors := make(chan error)

	for _, f := range filenames {
		go func(f string) {
			_, err := thumbnail.ImageFile(f)
			errors <- err
		}(f)
	}

	for range filenames {
		if err := <-errors; err != nil {
			return err // NOTE: incorrect: goroutine leak! // 可能存在goroutine泄漏!
		}
	}

	return nil
}

// 最简单的解决办法就是用一个具有合适大小的buffered channel，这样这些worker goroutine向channel中发送错误时就不会被阻塞。
// (一个可选的解决办法是创建一个另外的goroutine，当main goroutine返回第一个错误的同时去排空channel)
// makeThumbnails5 makes thumbnails for the specified files in parallel.
// It returns the generated file names in an arbitrary order,
// or an error if any step failed.
func makeThumbnails5(filenames []string) (thumbfiles []string, err error) {
	type item struct {
		thumbfile string
		err       error
	}

	ch := make(chan item, len(filenames))
	for _, f := range filenames {
		go func(f string) {
			var it item
			it.thumbfile, it.err = thumbnail.ImageFile(f)
			ch <- it
		}(f)
	}

	for range filenames {
		it := <-ch
		if it.err != nil {
			return nil, it.err
		}
		thumbfiles = append(thumbfiles, it.thumbfile)
	}

	return thumbfiles, nil
}

// 我们最后一个版本的makeThumbnails返回了新文件们的大小总计数(bytes)。
// 和前面的版本都不一样的一点是我们在这个版本里没有把文件名放在slice里，而是通过一个string的channel传过来，所以我们无法对循环的次数进行预测。
// 为了知道最后一个goroutine什么时候结束(最后一个结束并不一定是最后一个开始)，我们需要一个递增的计数器，在每一个goroutine启动时加一，在goroutine退出时减一。
// 这需要一种特殊的计数器，这个计数器需要在多个goroutine操作时做到安全并且提供在其减为零之前一直等待的一种方法。
// 注意Add和Done方法的不对称。Add是为计数器加一，必须在worker goroutine开始之前调用，而不是在goroutine中；否则的话我们没办法确定Add是在"closer" goroutine调用Wait之前被调用。
// 并且Add还有一个参数，但Done却没有任何参数；其实它和Add(-1)是等价的。
// 我们使用defer来确保计数器即使是在出错的情况下依然能够正确地被减掉。
// 上面的程序代码结构是当我们使用并发循环，但又不知道迭代次数时很通常而且很地道的写法。
// sizes channel携带了每一个文件的大小到main goroutine，在main goroutine中使用了range loop来计算总和。
// *重点:考虑一下另一种方案：如果等待操作被放在了main goroutine中，在循环之前，这样的话就永远都不会结束了,如果在循环之后，那么又变成了不可达的部分，因为没有任何东西去关闭这个channel，这个循环就永远都不会终止。
// *图8.5
// makeThumbnails6 makes thumbnails for each file received from the channel.
// It returns the number of bytes occupied by the files it creates.
func makeThumbnails6(filenames <-chan string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup // number of working goroutines
	for f := range filenames {
		wg.Add(1)
		// worker
		go func(f string) {
			defer wg.Done()
			thumb, err := thumbnail.ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(thumb) // OK to ignore error
			sizes <- info.Size()
		}(f)
	}

	// closer
	go func() {
		wg.Wait()
		close(sizes) // 关闭的channel依然可以进行接受已经成功发送的数据
	}()

	var total int64
	for size := range sizes { // 当sizes关闭且读取不到数据时,会退出range循环!(见pipeline2.go)
		total += size
	}

	return total
}
