package chapter9_7

import (
	"io/ioutil"
	"net/http"
)

// 并发的非阻塞缓存
// 本节中我们会做一个无阻塞的缓存，这种工具可以帮助我们来解决现实世界中并发程序出现但没有现成的库可以解决的问题。
// 这个问题叫作缓存(memoizing)函数，也就是说，我们需要缓存函数的返回结果，这样在对函数进行调用的时候，我们就只需要一次计算，之后只要返回计算的结果就可以了。
// 我们的解决方案会是并发安全且会避免对整个缓存加锁而导致所有操作都去争一个锁的设计。

// 我们将使用下面的httpGetBody函数作为我们需要缓存的函数的一个样例。
// 这个函数会去进行HTTP GET请求并且获取http响应body。对这个函数的调用本身开销是比较大的，所以我们尽量避免在不必要的时候反复调用。
func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func incomingURLs() []string {
	var urls []string = make([]string, 0)
	urls = append(urls, "http://www.baidu.com")
	urls = append(urls, "http://www.baidu.com") // 相同的url用于测试函数结果缓存效果
	urls = append(urls, "http://www.google.com")
	urls = append(urls, "http://www.bilibili.com")
	return urls
}

// 包级类型定义:
// Func is the type of the function to memoize.
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}
