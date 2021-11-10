package chapter5_2

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

// !!当outline调用自身时，被调用者接收的是stack的拷贝。
// 被调用者对stack的元素追加操作，修改的是stack的拷贝,但这个过程并不会修改调用方的stack。(即:父节点和每一个子节点拼接)
func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // push tag
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

func Outline() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	outline(nil, doc)
}
