package chapter5_5

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

// usage :
// 1. 编译chapter1_5.FetchUrl()
// 2. 执行./exercise http://www.baidu.com  > data/content
// 3. 编译chapter5_2.Outline()
// 4. 执行./exercise < data/content

// forEachNode针对每个结点x,都会调用pre(x)和post(x)。
// pre和post都是可选的。
// 遍历孩子结点之前,pre被调用
// 遍历孩子结点之后，post被调用
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		// %*s中的*会在字符串之前填充一些空格。
		// 在例子中，每次输出会先填充depth*2数量的空格，再输出""，最后再输出HTML标签。
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}

// 与之前的outline程序相比，我们得到了更加详细的页面结构
func Outline2() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	forEachNode(doc, startElement, endElement)
}
