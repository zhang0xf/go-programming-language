package chapter7_5

import (
	"bytes"
	"fmt"
	"io"
)

// 一个不包含任何值的nil接口值和一个刚好包含nil指针的接口值是不同的。这个细微区别产生了一个容易绊倒每个Go程序员的陷阱。

const debug = true

func InterfaceValueAndNil() {
	var buf *bytes.Buffer // struct
	fmt.Printf("%T\n", buf)
	if debug {
		buf = new(bytes.Buffer) // enable collection of output
	}
	f(buf) // NOTE: subtly incorrect!
	if debug {
		// ...use buf...
	}
}

// If out is non-nil, output will be written to it.
func f(out io.Writer) {
	fmt.Printf("%T\n", out)
	// ...do something...
	if out != nil {
		out.Write([]byte("done!\n")) // panic: nil pointer dereference(when debug = false)
	}
}

// 我们可能会预计当把变量debug设置为false时可以禁止对输出的收集，但是实际上在out.Write方法调用时程序发生了panic：
