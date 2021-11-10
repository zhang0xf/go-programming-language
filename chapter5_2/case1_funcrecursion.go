package chapter5_2

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

// 函数递归

// download or install golang.org/x/net :
// 1. go get -u golang.org/x/net
// 2. go build

// usage :
// 1. 编译chapter1_5.FetchUrl()
// 2. 执行./exercise http://www.baidu.com  > data/content
// 3. 编译chapter5_1.FuncRecursion()
// 4. 执行./exercise < data/content

// visit appends to links each link found in n and returns the result.
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

// 函数解析HTML标准输入，通过递归函数visit获得links（链接），并打印出这些links：
func FindLinks1() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

// 大部分编程语言使用固定大小的函数调用栈，常见的大小从64KB到2MB不等。
// 固定大小栈会限制递归的深度，当你用递归处理大量数据时，需要避免栈溢出；除此之外，还会导致安全性问题。
// 与此相反，Go语言使用可变栈，栈的大小按需增加(初始时很小)。这使得我们使用递归时不必考虑溢出和安全问题。
