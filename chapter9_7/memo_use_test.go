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

// 下一个Get的实现，调用Get的goroutine会两次获取锁：查找阶段获取一次，如果查找没有返回任何内容，那么进入更新阶段会再次获取。
// 在这两次获取锁的中间阶段，其它goroutine可以随意使用cache。
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
