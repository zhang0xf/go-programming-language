package chapter9_7

import (
	"fmt"
	"log"
	"sync"
	"testing"
	"time"
)

// usage : go test -v ./chapter9_7/
//         go test -run=TestMemoUse1 -v ./chapter9_7
//         go test -run=TestMemoUse2 -race -v ./chapter9_7/

// 尽管我们第一次对每一个URL的(*Memo).Get的调用都会花上55.438415ms毫秒，但第二次就只需要花256ns纳秒就可以返回完整的数据了。

// 这个测试是顺序地去做所有的调用的。
func TestMemoUse(t *testing.T) {
	m := New1(httpGetBody)
	for _, url := range incomingURLs() {
		start := time.Now()
		value, err := m.Get(url)
		if err != nil {
			log.Print(err)
		}
		fmt.Printf("%s, %s, %d bytes\n",
			url, time.Since(start), len(value.([]byte)))
	}
}

// 由于这种彼此独立的HTTP请求可以很好地并发，我们可以把这个测试改成并发形式。可以使用sync.WaitGroup来等待所有的请求都完成之后再返回。
// 这次测试跑起来更快了，然而不幸的是貌似这个测试不是每次都能够正常工作。
// 我们注意到有一些意料之外的cache miss(缓存未命中)，或者命中了缓存但却返回了错误的值，或者甚至会直接崩溃。
// 但更糟糕的是，有时候这个程序还是能正确的运行，所以我们甚至可能都不会意识到这个程序有bug。
// 但是我们可以使用-race这个flag来运行程序，竞争检测器(§9.6)会打印报告：
// WARNING: DATA RACE
// ...
// 具体:memo.go的32行出现了两次，说明有两个goroutine在没有同步干预的情况下更新了cache map。这表明Get不是并发安全的，存在数据竞争。
// 28  func (memo *Memo) Get(key string) (interface{}, error) {
// 29      res, ok := memo.cache(key)
// 30      if !ok {
// 31          res.value, res.err = memo.f(key)
// 32          memo.cache[key] = res
// 33      }
// 34      return res.value, res.err
// 35  }
func TestMemoUse1(t *testing.T) {
	m := New1(httpGetBody)
	var n sync.WaitGroup
	for _, url := range incomingURLs() {
		n.Add(1)
		go func(url string) {
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				log.Print(err)
			}
			fmt.Printf("%s, %s, %d bytes\n",
				url, time.Since(start), len(value.([]byte)))
			n.Done()
		}(url)
	}
	n.Wait()
}

// 最简单的使cache并发安全的方式是使用基于监控的同步。
// 只要给Memo加上一个mutex，在Get的一开始获取互斥锁，return的时候释放锁，就可以让cache的操作发生在临界区内了：
// 测试依然并发进行，但这回竞争检查器“沉默”了。
// 不幸的是对于Memo的这一点改变使我们完全丧失了并发的性能优点。
// 每次对f的调用期间都会持有锁，Get将本来可以并行运行的I/O操作串行化了。
// 我们本章的目的是完成一个无锁缓存，而不是现在这样的将所有请求串行化的函数的缓存。
func TestMemoUse2(t *testing.T) {
	m := New2(httpGetBody)
	var n sync.WaitGroup
	for _, url := range incomingURLs() {
		n.Add(1)
		go func(url string) {
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				log.Print(err)
			}
			fmt.Printf("%s, %s, %d bytes\n",
				url, time.Since(start), len(value.([]byte)))
			n.Done()
		}(url)
	}
	n.Wait()
}

// 下一个Get的实现，调用Get的goroutine会两次获取锁：查找阶段获取一次，如果查找没有返回任何内容，那么进入更新阶段会再次获取。在这两次获取锁的中间阶段，其它goroutine可以随意使用cache。
// 这些修改使性能再次得到了提升，但有一些URL被获取了两次。这种情况在两个以上的goroutine同一时刻调用Get来请求同样的URL时会发生。
// 多个goroutine一起查询cache，发现没有值，然后一起调用f这个慢不拉叽的函数。在得到结果后，也都会去更新map。其中一个获得的结果会覆盖掉另一个的结果。
func TestMemoUse3(t *testing.T) {
	m := New3(httpGetBody)
	var n sync.WaitGroup
	for _, url := range incomingURLs() {
		n.Add(1)
		go func(url string) {
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				log.Print(err)
			}
			fmt.Printf("%s, %s, %d bytes\n",
				url, time.Since(start), len(value.([]byte)))
			n.Done()
		}(url)
	}
	n.Wait()
}

// 理想情况下是应该避免掉多余的工作的。而这种“避免”工作一般被称为duplicate suppression(重复抑制/避免)。
// 下面版本的Memo每一个map元素都是指向一个条目的指针。每一个条目包含对函数f调用结果的内容缓存。
// 与之前不同的是这次entry还包含了一个叫ready的channel。在条目的结果被设置之后，这个channel就会被关闭，以向其它goroutine广播(§8.9)去读取该条目内的结果是安全的了。
// 现在Get函数包括下面这些步骤了：获取互斥锁来保护共享变量cache map，查询map中是否存在指定条目，如果没有找到那么分配空间插入一个新条目，释放互斥锁。
// 如果存在条目的话且其值没有写入完成(也就是有其它的goroutine在调用f这个慢函数)时，goroutine必须等待值ready之后才能读到条目的结果。
// 而想知道是否ready的话，可以直接从ready channel中读取，由于这个读取操作在channel关闭之前一直是阻塞。
// 如果没有条目的话，需要向map中插入一个没有准备好的条目，当前正在调用的goroutine就需要负责调用慢函数、更新条目以及向其它所有goroutine广播条目已经ready可读的消息了。
// 条目中的e.res.value和e.res.err变量是在多个goroutine之间共享的。
// 创建条目的goroutine同时也会设置条目的值，其它goroutine在收到"ready"的广播消息之后立刻会去读取条目的值。尽管会被多个goroutine同时访问，但却并不需要互斥锁。
// ready channel的关闭一定会发生在其它goroutine接收到广播事件之前，因此第一个goroutine对这些变量的写操作是一定发生在这些读操作之前的。不会发生数据竞争。
// 这样并发、不重复、无阻塞的cache就完成了。
func TestMemoUse4(t *testing.T) {
	m := New4(httpGetBody)
	var n sync.WaitGroup
	for _, url := range incomingURLs() {
		n.Add(1)
		go func(url string) {
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				log.Print(err)
			}
			fmt.Printf("%s, %s, %d bytes\n",
				url, time.Since(start), len(value.([]byte)))
			n.Done()
		}(url)
	}
	n.Wait()
}
