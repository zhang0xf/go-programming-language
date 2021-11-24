package chapter7_3

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// 一个类型如果拥有一个接口需要的所有方法，那么这个类型就实现了这个接口。
// 例如，*os.File类型实现了io.Reader，Writer，Closer，和ReadWriter接口。
// 例如，*bytes.Buffer实现了Reader，Writer，和ReadWriter这些接口，但是它没有实现Closer接口因为它不具有Close方法。

// 接口指定的规则非常简单：表达一个类型属于某个接口只要这个类型实现这个接口。

func ImplementInterface() {
	var w io.Writer
	w = os.Stdout         // OK: *os.File has Write method
	w = new(bytes.Buffer) // OK: *bytes.Buffer has Write method
	// w = time.Second       // compile error: time.Duration lacks Write method

	var rwc io.ReadWriteCloser
	rwc = os.Stdout // OK: *os.File has Read, Write, Close methods
	// rwc = new(bytes.Buffer) // compile error: *bytes.Buffer lacks Close method

	fmt.Printf("%T", w)
	fmt.Printf("%T", rwc)

	w = rwc // OK: io.ReadWriteCloser has Write method
	// rwc = w // compile error: io.Writer lacks Close method

	// 因为ReadWriter和ReadWriteCloser包含所有Writer的方法，所以任何实现了ReadWriter和ReadWriteCloser的类型必定也实现了Writer接口
}
